package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单完结 APIResponse
tmall.servicecenter.workcard.complete

工单完结
*/
type TmallServicecenterWorkcardCompleteAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardCompleteResponse
}

type TmallServicecenterWorkcardCompleteResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_complete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
