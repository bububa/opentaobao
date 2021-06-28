package util

// BaseRep 
/* model for simplify = false
type BaseRep struct {

    // 返回错误消息
    
    Msg   string `json:"msg,omitempty"`
    

    // 返回错误code
    
    Code   string `json:"code,omitempty"`
    

    // 返回结果
    
    Data   bool `json:"data,omitempty"`
    

    // 内层大对象
    
    Datas  struct {
        AssetQrCodeDto  []AssetQrCodeDto `json:"asset_qr_code_dto,omitempty"`
    } `json:"datas,omitempty"`
    

}
*/

// BaseRep 
type BaseRep struct {

    // 返回错误消息
    Msg   string `json:"msg,omitempty"`

    // 返回错误code
    Code   string `json:"code,omitempty"`

    // 返回结果
    Data   bool `json:"data,omitempty"`

    // 内层大对象
    Datas   []AssetQrCodeDto `json:"datas,omitempty"`

}
