package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商分派工人 APIResponse
tmall.servicecenter.workcard.assignworker

服务商调用该接口分派工人给具体的工单
*/
type TmallServicecenterWorkcardAssignworkerAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardAssignworkerResponse
}

type TmallServicecenterWorkcardAssignworkerResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_assignworker_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // -
    
    Result   *TmallServicecenterWorkcardAssignworkerResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
