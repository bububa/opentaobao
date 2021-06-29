package xhotelonlineorder

// TopInvoiceDO 
type TopInvoiceDO struct {
    // 发票提供方:0未知 1酒店前台，2卖家开具
    ProviderType   int64 `json:"provider_type,omitempty" xml:"provider_type,omitempty"`
}
