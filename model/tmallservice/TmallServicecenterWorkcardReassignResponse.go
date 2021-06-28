package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单改派门店 APIResponse
tmall.servicecenter.workcard.reassign

工单改派门店
*/
type TmallServicecenterWorkcardReassignAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardReassignResponse
}

type TmallServicecenterWorkcardReassignResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_reassign_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *TmallServicecenterWorkcardReassignResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
