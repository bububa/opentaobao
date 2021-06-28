package ma

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建码平台常用二维码 APIResponse
taobao.ma.qrcode.common.create

创建码平台对外提供的常用二维码接口，不适于码平台业务类型的码创建，如不支持包裹码、媒体码等，业务类型的码需要单独提供API。
*/
type TaobaoMaQrcodeCommonCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMaQrcodeCommonCreateResponse `json:"ma_qrcode_common_create_response,omitempty"` 
    TaobaoMaQrcodeCommonCreateResponse
}

/* model for simplify = false
type TaobaoMaQrcodeCommonCreateResponse struct {

    // 二维码对像
    
    Modules  struct {
        QrcodeDO  []QrcodeDO `json:"qrcode_do,omitempty"`
    } `json:"modules,omitempty"`
    

    // 执行是否成功
    
    Suc   bool `json:"suc,omitempty"`
    

}
*/

type TaobaoMaQrcodeCommonCreateResponse struct {

    // 二维码对像
    Modules   []QrcodeDO `json:"modules,omitempty"`

    // 执行是否成功
    Suc   bool `json:"suc,omitempty"`

}
