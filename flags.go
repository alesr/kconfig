package kconfig

import (
	"flag"
	"log"
	"os"
)

func setFlags() {

	printFlagPtr := flag.Bool("p", false, "Print JSON configuration file.")
	outputFlagPtr := flag.Bool("o", true, "Output Braintree mode.")
	modifyFlagPtr := flag.Bool("m", false, "Modify Braintree mode.")

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

	if *outputFlagPtr {
		_, fileExist, err := checkFile()
		if err != nil {
			log.Fatal(err)
		}

		if !fileExist {
			ask4file()
		} else {
			c, err := decode()
			if err != nil {
				log.Fatal(err)
			}
			c.checkoutMode()
		}
	}
	os.Exit(0)
}
