package cainiaoncwl

// JhAreaInfo 
type JhAreaInfo struct {
    // 可以指定某个省的集货单
    Provice   string `json:"provice,omitempty" xml:"provice,omitempty"`
    // 可以指定某个省市的集货单
    District   string `json:"district,omitempty" xml:"district,omitempty"`
    // 可以指定省市区的集货单
    City   string `json:"city,omitempty" xml:"city,omitempty"`
}
