package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建核销单 APIResponse
alibaba.servicecenter.identifytask.create

创建核销单
*/
type AlibabaServicecenterIdentifytaskCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_servicecenter_identifytask_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"