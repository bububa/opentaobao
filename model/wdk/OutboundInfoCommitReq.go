package wdk

// OutboundInfoCommitReq 
/* model for simplify = false
type OutboundInfoCommitReq struct {

    // 商品列表
    
    OutboundItemInfos  struct {
        OutboundItemInfo  []OutboundItemInfo `json:"outbound_item_info,omitempty"`
    } `json:"outbound_item_infos,omitempty"`
    

    // 预计到货时间
    
    EstimatedArrivalAt   string `json:"estimated_arrival_at,omitempty"`
    

    // 出货时间
    
    OutboundAt   string `json:"outbound_at,omitempty"`
    

    // 收货类型(信任收货、非信任收货)
    
    TrustedInbound   bool `json:"trusted_inbound,omitempty"`
    

    // 供应商名称
    
    SupplierName   string `json:"supplier_name,omitempty"`
    

    // 供应商编码
    
    SupplierCode   string `json:"supplier_code,omitempty"`
    

    // asn单号
    
    AsnOrderNo   string `json:"asn_order_no,omitempty"`
    

    // 商家编码，已经废弃
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

}
*/

// OutboundInfoCommitReq 
type OutboundInfoCommitReq struct {

    // 商品列表
    OutboundItemInfos   []OutboundItemInfo `json:"outbound_item_infos,omitempty"`

    // 预计到货时间
    EstimatedArrivalAt   string `json:"estimated_arrival_at,omitempty"`

    // 出货时间
    OutboundAt   string `json:"outbound_at,omitempty"`

    // 收货类型(信任收货、非信任收货)
    TrustedInbound   bool `json:"trusted_inbound,omitempty"`

    // 供应商名称
    SupplierName   string `json:"supplier_name,omitempty"`

    // 供应商编码
    SupplierCode   string `json:"supplier_code,omitempty"`

    // asn单号
    AsnOrderNo   string `json:"asn_order_no,omitempty"`

    // 商家编码，已经废弃
    MerchantCode   string `json:"merchant_code,omitempty"`

}
