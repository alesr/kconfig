package kconfig

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func setFlags() {

	printFlagPtr := flag.Bool("p", false, "Print JSON configuration file.")
	outputFlagPtr := flag.Bool("o", true, "Output Braintree mode.")
	modifyFlagPtr := flag.Bool("m", false, "Modify Braintree mode.")
	removeFlagPtr := flag.Bool("r", false, "Remove kconfig.json file from the system.")

	flag.Parse()

	if *printFlagPtr {
		*outputFlagPtr = false
		print()
	}

	if *modifyFlagPtr {
		*outputFlagPtr = false
		c, err := decode()
		if err != nil {
			log.Fatal(err)
		}
		c.changeCheckoutMode()
	}

	if *removeFlagPtr {
		*outputFlagPtr = false
		msg, err := remove()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s\n\n", msg)
	}

	if *outputFlagPtr {
		c, err := decode()
		if err != nil {
			log.Fatal(err)
		}
		c.checkoutMode()
	}

	os.Exit(0)
}
