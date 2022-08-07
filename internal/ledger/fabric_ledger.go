package ledger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	channelName      = "mychannel"
	contractType     = "basic"
	walletPath       = "wallet"
	walletLabel      = "appUser"
	org1MSPid        = "Org1MSP"
	createTXFuncName = "CreateTX"
)

var (
	orgConfigPath = filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
	)
	org1CCPPath = filepath.Join(
		orgConfigPath,
		"org1.example.com",
		"connection-org1.yaml",
	)
	org1CREDPath = filepath.Join(
		orgConfigPath,
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)
	org1CertPath = filepath.Join(org1CREDPath, "signcerts", "cert.pem")
	org1KeyDir   = filepath.Join(org1CREDPath, "keystore")
)

func init() {
	log.Println("PWD: ", os.Getenv("PWD"))
	log.Println("============ application-golang starts ============")
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	}
	err = os.RemoveAll(walletPath)
	if err != nil {
		log.Fatalf("Error removing wallet directory: %v", err)
	}
}

type Controller struct {
	gw *gateway.Gateway
	ct *gateway.Contract
}

// NewController starts a new service instance
func NewController() *Controller {
	service := new(Controller)
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}
	if !wallet.Exists(walletLabel) {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(org1CCPPath))),
		gateway.WithIdentity(wallet, walletLabel),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	service.gw = gw
	network, err := gw.GetNetwork(channelName)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}
	contract := network.GetContract(contractType)
	service.ct = contract
	// log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	// result, err := contract.SubmitTransaction("InitLedger")
	// if err != nil {
	// 	log.Fatalf("Failed to Submit transaction: %v", err)
	// }
	// log.Println(string(result))
	return service
}

func (s *Controller) Close() {
	s.gw.Close()
}

func (s *Controller) SubmitTx(binding string, timestamp int64) error {
	// log.Println("--> Submit Transaction: Invoke, function that adds a new asset")
	_, err := s.ct.SubmitTransaction(createTXFuncName, binding, strconv.FormatInt(timestamp, 10))
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
		return err
	}
	// log.Println(string(result))
	return nil
}

func (s *Controller) GetAllTXs() error {
	result, err := s.ct.EvaluateTransaction("GetAllTXs")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))
	return nil
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")

	// read the certificate pem
	cert, err := os.ReadFile(filepath.Clean(org1CertPath))
	if err != nil {
		return err
	}

	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(org1KeyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(org1KeyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(org1MSPid, string(cert), string(key))

	return wallet.Put(walletLabel, identity)
}
