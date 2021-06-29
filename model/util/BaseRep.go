package util

// BaseRep 
type BaseRep struct {
    // 返回错误消息
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 返回错误code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 返回结果
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`
    // 内层大对象
    Datas   []AssetQrCodeDto `json:"datas,omitempty" xml:"datas>asset_qr_code_dto,omitempty"`
}
