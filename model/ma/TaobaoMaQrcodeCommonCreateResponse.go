package ma

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建码平台常用二维码 APIResponse
taobao.ma.qrcode.common.create

创建码平台对外提供的常用二维码接口，不适于码平台业务类型的码创建，如不支持包裹码、媒体码等，业务类型的码需要单独提供API。
*/
type TaobaoMaQrcodeCommonCreateAPIResponse struct {
    model.CommonResponse
    TaobaoMaQrcodeCommonCreateResponse
}

type TaobaoMaQrcodeCommonCreateResponse struct {
    XMLName xml.Name `xml:"ma_qrcode_common_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 二维码对像
    
    Modules   []QrcodeDO `json:"modules,omitempty" xml:"modules>qrcode_do,omitempty"`
    
    
    // 执行是否成功
    
    Suc   bool `json:"suc,omitempty" xml:"suc,omitempty"`

    
}
