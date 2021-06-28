package simba

// ItemWordPackageDto 
/* model for simplify = false
type ItemWordPackageDto struct {

    // 词包id（1-流量智选，2-捡漏词包）
    
    WordPackageId   int64 `json:"word_package_id,omitempty"`
    

    // 开/关词包（0-关闭，1-开启）
    
    OnlineStatus   int64 `json:"online_status,omitempty"`
    

    // pc端出价（单位为分）
    
    PcBidPrice   int64 `json:"pc_bid_price,omitempty"`
    

    // 词包类型（1-流量智选， 19-捡漏词包）
    
    PackageType   int64 `json:"package_type,omitempty"`
    

    // 词包名称（可选：流量智选词包、捡漏词包）
    
    WordPackageName   string `json:"word_package_name,omitempty"`
    

    // 无线端出价（单位为分）
    
    WlBidPrice   int64 `json:"wl_bid_price,omitempty"`
    

}
*/

// ItemWordPackageDto 
type ItemWordPackageDto struct {

    // 词包id（1-流量智选，2-捡漏词包）
    WordPackageId   int64 `json:"word_package_id,omitempty"`

    // 开/关词包（0-关闭，1-开启）
    OnlineStatus   int64 `json:"online_status,omitempty"`

    // pc端出价（单位为分）
    PcBidPrice   int64 `json:"pc_bid_price,omitempty"`

    // 词包类型（1-流量智选， 19-捡漏词包）
    PackageType   int64 `json:"package_type,omitempty"`

    // 词包名称（可选：流量智选词包、捡漏词包）
    WordPackageName   string `json:"word_package_name,omitempty"`

    // 无线端出价（单位为分）
    WlBidPrice   int64 `json:"wl_bid_price,omitempty"`

}
