package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务核销 APIResponse
tmall.service.code.consume

天猫服务平台－服务核销
*/
type TmallServiceCodeConsumeAPIResponse struct {
    model.CommonResponse
    Response *TmallServiceCodeConsumeResponse `json:"tmall_service_code_consume_response,omitempty"`
}

type TmallServiceCodeConsumeResponse struct {

    // 服务工单ID
    Result   int64 `json:"result,omitempty"`

}
