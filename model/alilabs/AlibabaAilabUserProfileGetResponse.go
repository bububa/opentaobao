package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户信息 APIResponse
alibaba.ailab.user.profile.get

提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
*/
type AlibabaAilabUserProfileGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabUserProfileGetResponse
}

type AlibabaAilabUserProfileGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailab_user_profile_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口model
    
    Result   *AlibabaAilabUserProfileGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
