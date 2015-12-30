package kconfig

import (
	"flag"
	"fmt"
	"log"
)

func setFlags() {

	// Define flags.
	printFlagPtr := flag.Bool("p", false, "Print JSON configuration file.")
	outputFlagPtr := flag.Bool("o", true, "Output Braintree mode.")
	modifyFlagPtr := flag.Bool("m", false, "Modify Braintree mode.")
	removeFlagPtr := flag.Bool("r", false, "Remove kconfig.json file from the system.")

	flag.Parse()

	// To print kconfig.json content.
	if *printFlagPtr {
		*outputFlagPtr = false
		print()
	}

	// Open option to change the checkout's mode field.
	if *modifyFlagPtr {
		*outputFlagPtr = false
		c, err := decode()
		if err != nil {
			log.Fatal(err)
		}
		c.changeCheckoutMode()
	}

	// To remove existing kconfig.json.
	if *removeFlagPtr {
		*outputFlagPtr = false
		msg, err := remove()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s\n\n", msg)
	}

	// Output actual checkout mode.
	if *outputFlagPtr {
		c, err := decode()
		if err != nil {
			log.Fatal(err)
		}
		c.checkoutMode()
	}
}
