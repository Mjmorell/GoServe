package goserve

import "fmt"

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
	return
}

func example() {
	/*
		This shows the new method vs the old depracated method.
		The old method was 'easier' to program as was based on structures,
		but required a unique funtion and struct be defined for each table type

		The main goal for the restructure is to be easier for the developer,
		as long as the field names are known, information can be easily pulled/pushed.
		Without the restructure, this api would need to be configured to include
		the specific struct before data can be accessed. Plus a function to pull
		information would need to be created. None of this is difficult, but
		if for whatever reason this api is slowed in development, a struct may not
		be added for a user to use. This way, the api is as simple as possible.panic

		Here's to better Service-Now handles!
	*/

	client := Client{Username: "user", Password: "password", Instance: "instance.service-now.com/"}

	//New Method:

	original, num := client.PULL("u_computer_support", Filter("number")+IS("CS0012345"))
	fmt.Printf("The number of tickets returned by PULL is %d", num)
	pushable := make(map[string]string)
	pushable["description"] = original[0]["description"]
	pushable["assigned_to"] = original[0]["assigned_to"]
	client.PUSH("u_computer_support", Filter("number")+IS("CS0012347"), pushable)

	//Old Method (depracated, but still functional for older developments):

	originalDepracated := client.FilterRequests("u_computer_support", Filter("number")+IS("CS0012345"))
	var pushableDepracated Request
	pushableDepracated.Description = originalDepracated[0].Description
	pushableDepracated.AssignedTo = originalDepracated[0].AssignedTo
	client.Update("u_computer_support", Filter("number")+IS("CS0012347"), pushableDepracated)
}
