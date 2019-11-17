package main

import (
	"Zoho-Integration-/zohoapifuncs"
	"Zoho-Integration-/zohoapifuncs/datastore"
	"Zoho-Integration-/zohoapifuncs/structures"
	"log"
)

// working boilerplate for zoho integration
func main() {
	newlead := &structures.Lead{}
	// 1. converge your platform specific lead definition here
	// 2. then change the below function arguments to sync in with your "Lead" definitions..
	err := zohoapifuncs.GenerateLeadInZoho(newlead)
	if err != nil {
		log.Print("\n error while sending lead to zoho. Error :  ", err)
		dberr := datastore.AddToFailedLeadsDb(newlead) // add to db
		if dberr != nil {
			log.Print("\n Error while adding to failedleadsdb.. Error is", err, " and value failed is  ", newlead)
		}
	}

}
