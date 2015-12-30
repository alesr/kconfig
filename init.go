package kconfig

const filename = "kconfig.json"

// Config will be used to create a template JSON file
// containing some variables used to website's configuration.
type Config struct {
	Filename string     `json:"filename"` // "kconfig.json"
	filepath string     // Path to the directory where the program is executed.
	Products []*Product `json:"products"` // List of all products.
	Checkout *Checkout  `json:"checkout"` // Environment variables related to the checkout functioning.
}

// Checkout setup.
type Checkout struct {
	Mode      bool       `json:"mode"`      // Related to Braintree mode. False to Sandbox, true for Production.
	MaxQtd    byte       `json:"max_qtd"`   // Maximum quantity of items available to buy at once.
	Braintree *Braintree `json:"braintree"` // Braintree account credentials.
}

// Braintree have two distinct modes:
// Production and sandbox (tests).
type Braintree struct {
	Production *Setup `json:"production"`
	Sandbox    *Setup `json:"sandbox"`
}

// Setup Braintree account credentials.
// Data provided by Braintree after account creation.
type Setup struct {
	Environment string `json:"environment"`
	MerchantID  string `json:"merchant_id"`
	PublicKey   string `json:"public_key"`
	PrivateKey  string `json:"private_key"`
}

// Product specifications.
type Product struct {
	Name        string `json:"name"`
	AltName     string `json:"alternative_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Init the program.
func Init() {
	setFlags()
}
