package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户appkey授权 APIResponse
alibaba.idle.user.permit

用于记录登录用户与服务商的绑定关系，用于业务数据分发和授权校验
*/
type AlibabaIdleUserPermitAPIResponse struct {
    model.CommonResponse
    AlibabaIdleUserPermitResponse
}

type AlibabaIdleUserPermitResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_user_permit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
