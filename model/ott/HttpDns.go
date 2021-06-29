package ott

// HttpDns 
type HttpDns struct {
    // method
    Method   string `json:"method,omitempty" xml:"method,omitempty"`
    // tvHost
    TvHost   string `json:"tv_host,omitempty" xml:"tv_host,omitempty"`
    // dnsAddress
    DnsAddress   string `json:"dns_address,omitempty" xml:"dns_address,omitempty"`
}
