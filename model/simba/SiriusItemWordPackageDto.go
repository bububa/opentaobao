package simba

// SiriusItemWordPackageDto 
type SiriusItemWordPackageDto struct {

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 词包id
    WordPackageId   int64 `json:"word_package_id,omitempty"`

    // 词包在线状态（0:关闭，1:开启）
    OnlineStatus   int64 `json:"online_status,omitempty"`

    // pc端出价
    PcBidPrice   int64 `json:"pc_bid_price,omitempty"`

    // 词包类型（1:流量智选，19:捡漏词包）
    PackageType   int64 `json:"package_type,omitempty"`

    // 词包名称
    WordPackageName   string `json:"word_package_name,omitempty"`

    // 无线端出价
    WlBidPrice   int64 `json:"wl_bid_price,omitempty"`

}
