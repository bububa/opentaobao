package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV活动修改 APIResponse
alibaba.lsy.crm.activity.update

ISV活动修改
*/
type AlibabaLsyCrmActivityUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityUpdateResponse
}

type AlibabaLsyCrmActivityUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaLsyCrmActivityUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
