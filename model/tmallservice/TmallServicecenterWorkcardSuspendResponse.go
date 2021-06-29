package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单挂起 APIResponse
tmall.servicecenter.workcard.suspend

工单挂起
*/
type TmallServicecenterWorkcardSuspendAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardSuspendResponse
}

type TmallServicecenterWorkcardSuspendResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_suspend_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TmallServicecenterWorkcardSuspendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
