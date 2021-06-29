package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序获取登录用户信息 APIResponse
alibaba.lsy.miniapp.user.get

零售云小程序，通过授权码获取登录的卖家账号信息
*/
type AlibabaLsyMiniappUserGetAPIResponse struct {
    model.CommonResponse
    AlibabaLsyMiniappUserGetResponse
}

type AlibabaLsyMiniappUserGetResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_miniapp_user_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应内容
    
    Result   *MiniAppResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
