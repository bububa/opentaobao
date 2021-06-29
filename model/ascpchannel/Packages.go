package ascpchannel

// Packages 
type Packages struct {

    // 物流公司编码
    
    LogisticsCode   string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
    

    // 运单号
    
    ExpressCode   string `json:"express_code,omitempty" xml:"express_code,omitempty"`
    

    // 包裹编号
    
    PackageCode   string `json:"package_code,omitempty" xml:"package_code,omitempty"`
    

    // 包裹长度 (厘米)
    
    Length   string `json:"length,omitempty" xml:"length,omitempty"`
    

    // 包裹宽度 (厘米)
    
    Width   string `json:"width,omitempty" xml:"width,omitempty"`
    

    // 包裹高度 (厘米)
    
    Height   string `json:"height,omitempty" xml:"height,omitempty"`
    

    // 包裹重量 (千克)
    
    Weight   string `json:"weight,omitempty" xml:"weight,omitempty"`
    

    // 包裹体积 (升, L)
    
    Volume   string `json:"volume,omitempty" xml:"volume,omitempty"`
    

    // 商品信息
    
    Items   []Items `json:"items,omitempty" xml:"items,omitempty"`
    

}
