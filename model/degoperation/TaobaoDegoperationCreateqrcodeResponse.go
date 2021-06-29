package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
中奖生成二维码 APIResponse
taobao.degoperation.createqrcode

用户中奖后，生成二维码图片链接
*/
type TaobaoDegoperationCreateqrcodeAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationCreateqrcodeResponse
}

type TaobaoDegoperationCreateqrcodeResponse struct {
    XMLName xml.Name `xml:"degoperation_createqrcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 二维码链接
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // realSubCode
    
    RealSubCode   string `json:"real_sub_code,omitempty" xml:"real_sub_code,omitempty"`

    
    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
