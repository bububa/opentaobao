package ott

import (
	"sync"
)

// HttpDns 结构体
type HttpDns struct {
	// method
	Method string `json:"method,omitempty" xml:"method,omitempty"`
	// tvHost
	TvHost string `json:"tv_host,omitempty" xml:"tv_host,omitempty"`
	// dnsAddress
	DnsAddress string `json:"dns_address,omitempty" xml:"dns_address,omitempty"`
}

var poolHttpDns = sync.Pool{
	New: func() any {
		return new(HttpDns)
	},
}

// GetHttpDns() 从对象池中获取HttpDns
func GetHttpDns() *HttpDns {
	return poolHttpDns.Get().(*HttpDns)
}

// ReleaseHttpDns 释放HttpDns
func ReleaseHttpDns(v *HttpDns) {
	v.Method = ""
	v.TvHost = ""
	v.DnsAddress = ""
	poolHttpDns.Put(v)
}
