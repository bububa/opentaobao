package mos

// CodeInfoDto 
/* model for simplify = false
type CodeInfoDto struct {

    // 商品信息
    
    GoodsList  struct {
        CodeGoodsDto  []CodeGoodsDto `json:"code_goods_dto,omitempty"`
    } `json:"goods_list,omitempty"`
    

    // 包裹信息
    
    PackageCode   string `json:"package_code,omitempty"`
    

    // 寄件信息
    
    SendInfo  *struct {
        DeliveryCustomDto  *DeliveryCustomDto `json:"delivery_custom_dto,omitempty"`
    } `json:"send_info,omitempty"`
    

}
*/

// CodeInfoDto 
type CodeInfoDto struct {

    // 商品信息
    GoodsList   []CodeGoodsDto `json:"goods_list,omitempty"`

    // 包裹信息
    PackageCode   string `json:"package_code,omitempty"`

    // 寄件信息
    SendInfo   *DeliveryCustomDto `json:"send_info,omitempty"`

}
