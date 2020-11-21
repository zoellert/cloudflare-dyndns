package services

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
)

var (
	API    *cloudflare.API
	zoneId string
)

func InitCloudflare() {
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_MAIL"))
	if err != nil {
		log.Printf("Cloudflare authentication failed for user %s", os.Getenv("CLOUDFLARE_MAIL"))
		log.Fatal(err)
	}

	API = api

	id, err := api.ZoneIDByName(os.Getenv("CLOUDFLARE_ZONE"))
	if err != nil {
		log.Print("Failed to get Cloudflare zone id")
		log.Fatal(err)
	}

	zoneId = id
}

// Returns true if the ip has changed
// Returns false if the ip has not changed
func UpdateDNSRecord(name, ip string) (error, bool) {
	recordFilter := cloudflare.DNSRecord{Type: "A", Name: name}
	recs, err := API.DNSRecords(zoneId, recordFilter)
	if err != nil {
		log.Print("Failed to get Cloudflare dns records")
		log.Print(err)
		return fmt.Errorf("failed to get cloudflare dns records"), false
	}

	newRecord := cloudflare.DNSRecord{
		Type:      "A",
		Name:      name,
		Content:   ip,
		Proxiable: false,
		Proxied:   false,
	}

	if len(recs) < 1 {
		_, err := API.CreateDNSRecord(zoneId, newRecord)
		if err != nil {
			log.Printf("Failed to create dns record")
			log.Print(err)
			return fmt.Errorf("failed to create dns record"), false
		}
		return nil, true
	}

	for _, rec := range recs {
		if rec.Content == ip {
			return nil, false
		} else {
			err := API.UpdateDNSRecord(zoneId, rec.ID, newRecord)
			if err != nil {
				log.Printf("Failed to update dns record")
				log.Print(err)
				return fmt.Errorf("failed to udpate dns record"), false
			}
			return nil, true
		}
	}
	return fmt.Errorf("internal server error"), false
}
