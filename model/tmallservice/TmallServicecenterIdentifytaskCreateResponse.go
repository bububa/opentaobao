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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_identifytask_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // -
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"