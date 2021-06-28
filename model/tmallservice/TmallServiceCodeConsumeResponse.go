package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务核销 APIResponse
tmall.service.code.consume

天猫服务平台－服务核销
*/
type TmallServiceCodeConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_service_code_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 服务工单ID
    
    Result   int64 `json:"result,omitempty" xml:"