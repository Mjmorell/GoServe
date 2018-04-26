package goserve

import (
	"fmt"
)

/************************************************
	MIT License
	Details viewable in the Github Directory
	Copyright (c) 2018 Michael Morell
*************************************************
	Created by Michael Morell
	Released NULL
	Github: https://github.com/Mjmorell/GoServicenow
************************************************/


func main() {
	fmt.Println("tested!")
	//username, instance, password := login()
	//c := Client{
	//	Username: username,
	//	Password: password,
	//	Instance: instance,
	//}

	/*
	for {
		serial := Query_Generic("Scan a serial")
		if serial == "" {
			continue
		}
		if Config.Regex.UseRegex {
			serial = strings.Split(serial, Config.Regex.RegexSplit)[regexSectionWanted]
		}

		fmt.Println(serial)

		// temporary validation struct to verify one and only one asset is returned for the passed serial
		var updateCheckStruct []Asset
		updateCheckStruct = c.PerformForAsset("alm_hardware", "serial_number="+serial)

		// Checking only one serial was matched.
		if len(updateCheckStruct) > 1 {
			panic("Too many serials returned for " + serial)
		} else if len(updateCheckStruct) < 1 {
			panic("No serials returned for " + serial)
		}

		// The main update itself
		c.PushFor("alm_hardware", "serial_number="+serial, pushedStruct)

		// Checking that install status is set properly, else retry pushing then throw an error
		if updateCheckStruct[0].InstallStatus != "1" {
			c.PushFor("alm_hardware", "serial_number="+serial, pushedStruct)
			updateCheckStruct = c.PerformForAsset("alm_hardware", "serial_number="+serial)
			if updateCheckStruct[0].InstallStatus != "1" {
				panic("ERROR WITH " + serial)
			}
		}

		// Checking that location is set properly, else retry pushing then throw an error
		if updateCheckStruct[0].Location != pushedStruct.Location {
			c.PushFor("alm_hardware", "serial_number="+serial, pushedStruct)
			updateCheckStruct = c.PerformForAsset("alm_hardware", "serial_number="+serial)
			if updateCheckStruct[0].Location != pushedStruct.Location {
				panic("ERROR WITH " + serial)
			}
		}
	}*/
}
