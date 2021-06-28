package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
合单生成核销单 APIResponse
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单
*/
type AlibabaServicecenterFulfiltaskCreateAPIResponse struct {
    model.CommonResponse
    AlibabaServicecenterFulfiltaskCreateResponse
}

type AlibabaServicecenterFulfiltaskCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_servicecenter_fulfiltask_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaServicecenterFulfiltaskCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
