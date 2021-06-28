package mos

// CodeInfoDto 
type CodeInfoDto struct {

    // 商品信息
    
    GoodsList   []CodeGoodsDto `json:"goods_list,omitempty" xml:"goods_list,omitempty"`
    

    // 包裹信息
    
    PackageCode   string `json:"package_code,omitempty" xml:"package_code,omitempty"`
    

    // 寄件信息
    
    SendInfo   *DeliveryCustomDto `json:"send_info,omitempty" xml:"send_info,omitempty"`
    

}
