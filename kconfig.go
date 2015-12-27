package kconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Write config struct into a JSON file.
func (c *Config) write() error {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := dir + string(os.PathSeparator) + filename

	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	jsonPtr := &json

	err = ioutil.WriteFile(path, *jsonPtr, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Ask user for input and set new mode.
func (c *Config) changeCheckoutMode() {

	c, err := decode()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n[1] PRODUCTION\n[2] SANDBOX\n\n$ ")

	var input int
	_, err = fmt.Scan(&input)
	if err != nil || (input != 1 && input != 2) {
		fmt.Println(errors.New("\nType 1 for Production or 2 for Sandbox"))
		c.changeCheckoutMode()
	}

	if input == 1 {
		c.Checkout.Mode = true
	} else {
		c.Checkout.Mode = false
	}

	c.write()
	c.checkoutMode()
}

// Stdout Braintree chechout mode.
func (c *Config) checkoutMode() {
	if c.Checkout.Mode {
		fmt.Printf("\n   Braintree set to PRODUCTION mode.\n")
	} else {
		fmt.Printf("\n   Braintree set to SANDBOX mode.\n")
	}
}

// Stdout JSON config file.
func print() {
	path, fileExist, err := checkFile()
	if err != nil {
		log.Fatal(err)
	}

	if !fileExist {
		ask4file()
	} else {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n\n%s", string(content))
	}

}

// Decode the configuration JSON file into a struct Config.
func decode() (*Config, error) {

	c := new(Config)

	path, fileExist, err := checkFile()
	if err != nil {
		return nil, err
	}

	if !fileExist {
		ask4file()
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(file).Decode(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Config) new() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(errors.New("Failed to get directory."))
	}

	*c = Config{
		Filename: filename,
		Filepath: dir,
		Products: []*Product{
			&Product{
				Name:        "",
				Description: "",
				Price:       0,
			},
		},
		Checkout: &Checkout{
			Mode:   false,
			MaxQtd: 0,
			Braintree: &Braintree{
				Production: &Setup{
					Environment: "",
					MerchantID:  "",
					PublicKey:   "",
					PrivateKey:  "",
				},
				Sandbox: &Setup{
					Environment: "",
					MerchantID:  "",
					PublicKey:   "",
					PrivateKey:  "",
				},
			},
		},
	}

	err = c.write()
	if err != nil {
		log.Fatal("Failed to write JSON file\n", err)
	}

	fmt.Printf("\n%s was successfully created", filename)
}

func checkFile() (string, bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false, err
	}

	path := dir + string(os.PathSeparator) + filename

	// Checks if file exist in directory.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", false, nil
	}

	return path, true, nil
}

func ask4file() {
	var input string
	fmt.Print("\nCreate kconfig.json? Y / N\n\n$ ")

	_, err := fmt.Scan(&input)

	input = strings.ToLower(input)

	if err != nil || (input != "y" && input != "n") {
		fmt.Println(errors.New("\nType Y for yes or N for no"))
		ask4file()
	}

	if input == "y" {
		c := new(Config)
		c.new()
	} else {
		os.Exit(0)
	}
}
