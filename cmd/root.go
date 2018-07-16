package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/ejcx/cf/lib"
	"github.com/spf13/cobra"
)

var cfgFile string

type Credentials struct {
	Email string
}

var RootCmd = &cobra.Command{
	Use:   "cf",
	Short: "A CLI for interacting with Cloudflare's V4 API",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func Main(cmd *cobra.Command, args []string, name string) {
	err := lib.DefaultCredentialProvider.ConfigureEnvironment()
	if err != nil {
		log.Fatalf("No set of credentials to use: %s", err)
	}

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal("Could not initialize api object: %s", err)
	}

	r, err := root(cmd, args, name, api)
	if err != nil {
		log.Fatalf("Could not make cloudflare request: %s", err)
	}
	buf, err := json.MarshalIndent(r, " ", "    ")
	if err != nil {
		log.Fatal("Could not make print resp: %s", err)
	}
	fmt.Println(string(buf))
}

func root(cmd *cobra.Command, args []string, name string, api *cloudflare.API) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	switch name {
	case "ListZones":
		if ZoneNameFilter != "" {
			resp, err = api.ListZones(ZoneNameFilter)
		} else {
			resp, err = api.ListZones()
		}
	case "DNSRecords":
		rec := cloudflare.DNSRecord{}
		if Type != "" {
			rec.Type = Type
		}
		if Name != "" {
			rec.Name = Name
		}
		if Content != "" {
			rec.Content = Content
		}
		resp, err = api.DNSRecords(ZoneID, rec)
	case "CreateDNSRecord":
		rec := cloudflare.DNSRecord{}
		if Type != "" {
			rec.Type = Type
		}
		if Name != "" {
			rec.Name = Name
		}
		if Content != "" {
			rec.Content = Content
		}
		if Priority > 0 {
			rec.Priority = Priority
		}
		if Ttl != 0 {
			rec.TTL = Ttl
		}
		resp, err = api.CreateDNSRecord(ZoneID, rec)
	case "DeleteDNSRecord":
		err = api.DeleteDNSRecord(ZoneID, RecordID)
		if err == nil {
			resp = map[string]interface{}{
				"Success": true,
			}
		}
	case "DeleteZone":
		resp, err = api.DeleteZone(ZoneID)
	case "DNSRecord":
		resp, err = api.DNSRecord(ZoneID, RecordID)
	case "ListAllRateLimits":
		resp, err = api.ListAllRateLimits(ZoneID)
	case "ListLoadBalancers":
		resp, err = api.ListLoadBalancers(ZoneID)
	case "ListOrganizationsBalancers":
		resp, _, err = api.ListOrganizations()
	case "ListPageRules":
		resp, err = api.ListPageRules(ZoneID)
	case "ListCustomCerts":
		resp, err = api.ListSSL(ZoneID)
	case "ListWAFPackages":
		resp, err = api.ListWAFPackages(ZoneID)
	case "ListWAFRules":
		resp, err = api.ListWAFRules(ZoneID, PackageID)
	case "EditZonePaused":
		z := cloudflare.ZoneOptions{
			Paused: &Paused,
		}
		resp, err = api.EditZone(ZoneID, z)
	case "EditZoneVanityNS":
		vns := strings.Split(VanityNS, ",")
		z := cloudflare.ZoneOptions{
			VanityNS: vns,
		}
		resp, err = api.EditZone(ZoneID, z)
	case "ListZoneLockdowns":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListZoneLockdowns(ZoneID, page)
	case "ListUserAgentRules":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListUserAgentRules(ZoneID, page)
	case "CreateZone":
		org := cloudflare.Organization{}
		if OrganizationID != "" {
			org.ID = OrganizationID
		}
		resp, err = api.CreateZone(Name, false, org)
	default:
		break
	}
	return resp, err
}
