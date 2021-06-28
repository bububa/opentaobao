package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商创建核销单 APIResponse
tmall.servicecenter.identifytask.create

服务商调用该接口进行创建核销单操作
*/
type TmallServicecenterIdentifytaskCreateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterIdentifytaskCreateResponse
}

type TmallServicecenterIdentifytaskCreateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_identifytask_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // -
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
