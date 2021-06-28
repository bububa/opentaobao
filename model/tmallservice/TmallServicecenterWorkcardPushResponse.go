package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送服务工单信息 APIResponse
tmall.servicecenter.workcard.push

服务商家推送工单信息到天猫。
*/
type TmallServicecenterWorkcardPushAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardPushResponse
}

type TmallServicecenterWorkcardPushResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_push_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    WorkcardResult   *ResultBase `json:"workcard_result,omitempty" xml:"workcard_result,omitempty"`

    
}
