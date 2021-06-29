package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当前授权用户手机号码 APIResponse
taobao.miniapp.user.phone.get

在商家应用中，获取当前授权用户手机号码
*/
type TaobaoMiniappUserPhoneGetAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappUserPhoneGetResponse
}

type TaobaoMiniappUserPhoneGetResponse struct {
    XMLName xml.Name `xml:"miniapp_user_phone_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 用户手机号码
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`

    
}
