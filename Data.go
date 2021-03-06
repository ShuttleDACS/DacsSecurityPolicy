package main

// the configuration object
type Configuration struct {
	PortNumber   	string
	WebServer    	string
	WalletServer 	string
	WalletConnected string
}

type StructKeys struct {
	PrivateKey string `bson:"_PrivateKey"`
	PublicKey  string `bson:"_PublicKey"`
}

type StructSendTransaction struct {
	PublicKey string `json:"publicKey"`
	Message   string `json:"message"`
}

type StructSendTransactionTest struct {
	PublicKey string `json:"publicKey"`
	Message   string `json:"message"`
	Nonce     string `json:"nonce"`
	Sender    string `json:"sender"`
}

type StructCreateSignedTransaction struct {
	PrivateKey string `json:"privateKey"`
	Message    string `json:"message"`
}

type StructSetVal struct {
	ID string `json:"id"`
}

type StuctWalletTransaction struct {
	Action string   `json:"action"`
	Params []string `json:"params"`
}
