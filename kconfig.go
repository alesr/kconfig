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

	// Get directory path.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// dir + filename
	path := dir + string(os.PathSeparator) + filename

	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	jsonPtr := &json

	// Writes json data to file.
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

	// Ask user for input mode.
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

	// Write the new mode to the file and output it.
	c.write()
	c.checkoutMode()
}

// Stdout Braintree chechout mode.
func (c *Config) checkoutMode() {
	if c.Checkout.Mode {
		fmt.Printf("\n   Braintree set to PRODUCTION mode\n\n")
	} else {
		fmt.Printf("\n   Braintree set to SANDBOX mode\n\n")
	}
}

// Stdout JSON config file.
func print() {
	path, exist, err := checkFile()
	if err != nil {
		log.Fatal(err)
	}

	// If file does not exist ask user to create one.
	if !exist {
		ask4file()
	} else {
		// Read it content and output to the console.
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

	path, exist, err := checkFile()
	if err != nil {
		return nil, err
	}

	if !exist {
		ask4file()
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		// Decode file content to Config struct.
		err = json.NewDecoder(file).Decode(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Create a new kconfig.json with zero values.
func newKconfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(errors.New("Failed to get directory."))
	}

	c := Config{
		Filename: filename,
		filepath: dir,
		Products: []*Product{},
		Checkout: &Checkout{
			Braintree: &Braintree{
				Production: &Setup{},
				Sandbox:    &Setup{},
			},
		},
	}

	err = c.write()
	if err != nil {
		log.Fatal("Failed to write JSON file\n", err)
	}
	fmt.Printf("\n%s was successfully created\n\n", filename)
}

// Check for a file returning its path and boolean referring to its existence.
func checkFile() (string, bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false, err
	}

	path := dir + string(os.PathSeparator) + filename

	// Check if file exist in directory.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", false, nil
	}

	return path, true, nil
}

// Ask user to create a new kconfig.json
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
		newKconfig()
	}
	os.Exit(0)
}

// Remove kconfig.json
func remove() (string, error) {

	path, exist, err := checkFile()
	if err != nil {
		return "", err
	}

	if !exist {
		return fmt.Sprintf("Can't remove %s. File does not exist.", filename), nil
	}

	// Ask user input.
	var input string
	fmt.Printf("\nAre you sure you want to remove %s? Y / N\n\n$ ", filename)
	_, err = fmt.Scan(&input)
	input = strings.ToLower(input)

	if err != nil || (input != "y" && input != "n") {
		fmt.Println(errors.New("\nType Y for yes or N for no"))
		remove()
	}

	if err := os.Remove(path); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s was successfully deleted", filename), nil
}
