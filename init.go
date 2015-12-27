package kconfig

const filename = "kconfig.json"

// Config will be used to create a JSON
// file for traditionalbox website's configuration.
type Config struct {
	Filename string     `json:"filename"`
	Filepath string     `json:"filepath"`
	Products []*Product `json:"products"`
	Checkout *Checkout  `json:"checkout"`
}

// Checkout setup
type Checkout struct {
	Mode      bool       `json:"mode"`
	MaxQtd    byte       `json:"max_qtd"`
	Braintree *Braintree `json:"braintree"`
}

// Braintree setup
type Braintree struct {
	Production *Setup `json:"production"`
	Sandbox    *Setup `json:"sandbox"`
}

// Setup Braintree account credentials
type Setup struct {
	Environment string `json:"environment"`
	MerchantID  string `json:"merchant_id"`
	PublicKey   string `json:"public_key"`
	PrivateKey  string `json:"private_key"`
}

// Product specifications
type Product struct {
	Name        string `json:"name"`
	AltName     string `json:"alternative_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Init the program
func Init() {
	setFlags()
}
