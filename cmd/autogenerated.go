package cmd

import "github.com/spf13/cobra"

var (
	ZoneNameFilter string
	ZoneID         string
	Type           string
	Name           string
	Content        string
	Ttl            int
	RecordID       string
	OrganizationID string
	Page           int
)

func init() {
	var ListZones = &cobra.Command{
		Use:   "list-zones",
		Short: "Command for listing zones",
		Long:  `  This is a meaty description of the list-zones`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListZones")
		},
	}

	ListZones.Flags().StringVar(&ZoneNameFilter, "zone-name-filter", "", "string for filtering by name")

	var ListDnsRecords = &cobra.Command{
		Use:   "list-dns-records",
		Short: "Command for listing dns-records",
		Long:  `  List DNS Records associated with a given zone-id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DNSRecords")
		},
	}

	ListDnsRecords.Flags().StringVar(&ZoneID, "zoneID", "", "zone id used for filtering")
	ListDnsRecords.MarkFlagRequired("zoneID")

	ListDnsRecords.Flags().StringVar(&Type, "type", "", "DNS Record type used for filter")

	ListDnsRecords.Flags().StringVar(&Name, "name", "", "DNS Record name used for filter")

	ListDnsRecords.Flags().StringVar(&Content, "content", "", "DNS Record content used for filter")

	var CreateDnsRecord = &cobra.Command{
		Use:   "create-dns-record",
		Short: "Command DNS Record",
		Long:  `Create DNS record associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateDNSRecord")
		},
	}

	CreateDnsRecord.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* zone id associated with the new dns record")
	CreateDnsRecord.MarkFlagRequired("zoneID")

	CreateDnsRecord.Flags().StringVar(&Type, "type", "", "*Required:* valid values: A, AAAA, CNAME, TXT, SRV, LOC, MX, NS, SPF, CERT, DNSKEY, DS, NAPTR, SMIMEA, SSHFP, TLSA, URI read only")
	CreateDnsRecord.MarkFlagRequired("type")

	CreateDnsRecord.Flags().StringVar(&Name, "name", "", "*Required:* DNS Record name (example: example.com), max length: 255")
	CreateDnsRecord.MarkFlagRequired("name")

	CreateDnsRecord.Flags().StringVar(&Content, "content", "", "*Required:* DNS Record content used for filter")
	CreateDnsRecord.MarkFlagRequired("content")

	CreateDnsRecord.Flags().IntVar(&Ttl, "ttl", 0, "Time to live for DNS record. Value of 1 is 'automatic', min value:120 max value:2147483647")

	var DeleteDnsRecord = &cobra.Command{
		Use:   "delete-dns-record",
		Short: "Delete DNS Record",
		Long:  `Delete DNS record associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteDNSRecord")
		},
	}

	DeleteDnsRecord.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* zone id associated with the record you wish to delete")
	DeleteDnsRecord.MarkFlagRequired("zoneID")

	DeleteDnsRecord.Flags().StringVar(&RecordID, "recordID", "", "*Required:* record id associated with the dns record you wish to delete")
	DeleteDnsRecord.MarkFlagRequired("recordID")

	var DeleteZone = &cobra.Command{
		Use:   "delete-zone",
		Short: "Delete zone",
		Long:  `Delete a zone associated with your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteZone")
		},
	}

	DeleteZone.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* zone id that will be deleted")
	DeleteZone.MarkFlagRequired("zoneID")

	var CreateZone = &cobra.Command{
		Use:   "create-zone",
		Short: "Create zone",
		Long:  `Create a zone associated with your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateZone")
		},
	}

	CreateZone.Flags().StringVar(&Name, "name", "", "*Required:* the zone name that will be added to your account")
	CreateZone.MarkFlagRequired("name")

	CreateZone.Flags().StringVar(&OrganizationID, "organizationID", "", "The organizationID associated with the zone")

	var ShowDnsRecord = &cobra.Command{
		Use:   "show-dns-record",
		Short: "Show DNS Record",
		Long:  `Show a single DNS record associated with a zone ID and record ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DNSRecord")
		},
	}

	ShowDnsRecord.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the DNS Record")
	ShowDnsRecord.MarkFlagRequired("zoneID")

	ShowDnsRecord.Flags().StringVar(&RecordID, "recordID", "", "*Reqiured:* The recordID associated with the DNS Record")
	ShowDnsRecord.MarkFlagRequired("recordID")

	var ListRatelimits = &cobra.Command{
		Use:   "list-ratelimits",
		Short: "Show Ratelimits",
		Long:  `Returns all Rate Limits for a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListAllRateLimits")
		},
	}

	ListRatelimits.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the Ratelimits")
	ListRatelimits.MarkFlagRequired("zoneID")

	var ListLoadbalancers = &cobra.Command{
		Use:   "list-loadbalancers",
		Short: "Show LoadBalancers",
		Long:  `Returns all LoadBalancers for a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListLoadBalancers")
		},
	}

	ListLoadbalancers.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the Ratelimits")
	ListLoadbalancers.MarkFlagRequired("zoneID")

	var ListOrganizations = &cobra.Command{
		Use:   "list-organizations",
		Short: "Show Organizations",
		Long:  `Returns all Organizations associated with your account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListOrganizations")
		},
	}

	var ListPageRules = &cobra.Command{
		Use:   "list-page-rules",
		Short: "Show Page Rules",
		Long:  `Returns all page rules associated with a given zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListPageRules")
		},
	}

	ListPageRules.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the pagerules")
	ListPageRules.MarkFlagRequired("zoneID")

	var ListCustomCerts = &cobra.Command{
		Use:   "list-custom-certs",
		Short: "Show Custom Certs",
		Long:  `Returns all custom certs for a given zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListCustomCerts")
		},
	}

	ListCustomCerts.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the custom certs")
	ListCustomCerts.MarkFlagRequired("zoneID")

	var ListUserAgentRules = &cobra.Command{
		Use:   "list-user-agent-rules",
		Short: "List User-Agent rules",
		Long:  `Returns all User-Agent rules for a specific zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListUserAgentRules")
		},
	}

	ListUserAgentRules.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the user-agent rule.")
	ListUserAgentRules.MarkFlagRequired("zoneID")

	ListUserAgentRules.Flags().IntVar(&Page, "page", 0, "*Required:* Pagination for user-agent rules")

	var ListWafPackages = &cobra.Command{
		Use:   "list-waf-packages",
		Short: "List WAF Packages",
		Long:  `Return the WAF Packages associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListWAFPackages")
		},
	}

	ListWafPackages.Flags().StringVar(&ZoneID, "zoneID", "", "*Required:* The zone ID associated with the WAF packages.")
	ListWafPackages.MarkFlagRequired("zoneID")

	var Zone = &cobra.Command{
		Use:   "zone",
		Short: "Commands for interacting with zones",
		Long:  `  This is a meaty description of the zone api.`,
	}
	Zone.AddCommand(ListZones)
	Zone.AddCommand(DeleteZone)
	Zone.AddCommand(CreateZone)

	RootCmd.AddCommand(Zone)

	var Dns = &cobra.Command{
		Use:   "dns",
		Short: "Commands for interacting with dns records",
		Long:  `  This is a meaty description of the dns api.`,
	}
	Dns.AddCommand(ListDnsRecords)
	Dns.AddCommand(ShowDnsRecord)
	Dns.AddCommand(CreateDnsRecord)
	Dns.AddCommand(DeleteDnsRecord)

	RootCmd.AddCommand(Dns)

	var Ssl = &cobra.Command{
		Use:   "ssl",
		Short: "Commands for interacting with ssl configuration",
		Long:  `  This is a meaty description of the ssl api.`,
	}
	Ssl.AddCommand(ListCustomCerts)

	RootCmd.AddCommand(Ssl)

	var Pagerules = &cobra.Command{
		Use:   "pagerules",
		Short: "Commands for interacting with pagerules api",
		Long:  `  This is a meaty description of the pagerules api.`,
	}
	Pagerules.AddCommand(ListPageRules)

	RootCmd.AddCommand(Pagerules)

	var Firewall = &cobra.Command{
		Use:   "firewall",
		Short: "Commands for interacting with firewall",
		Long:  `  This is a meaty description of the firewall apis.`,
	}
	Firewall.AddCommand(ListUserAgentRules)
	Firewall.AddCommand(ListWafPackages)

	RootCmd.AddCommand(Firewall)

	var Organization = &cobra.Command{
		Use:   "organization",
		Short: "Commands for interacting with organizations api",
		Long:  `  This is a meaty description of the organizaiton api.`,
	}
	Organization.AddCommand(ListOrganizations)

	RootCmd.AddCommand(Organization)

	var Ratelimit = &cobra.Command{
		Use:   "ratelimit",
		Short: "Commands for interacting with ratelimit api",
		Long:  `  This is a meaty description of the ratelimit api.`,
	}
	Ratelimit.AddCommand(ListRatelimits)

	RootCmd.AddCommand(Ratelimit)

	var Loadbalancer = &cobra.Command{
		Use:   "loadbalancer",
		Short: "Commands for interacting with loadbalancer api",
		Long:  `  This is a meaty description of the loadbalancer api.`,
	}
	Loadbalancer.AddCommand(ListLoadbalancers)

	RootCmd.AddCommand(Loadbalancer)

}