package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创单接口 APIResponse
alibaba.tcls.fulfill.qa.order.create

根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
*/
type AlibabaTclsFulfillQaOrderCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTclsFulfillQaOrderCreateResponse `json:"alibaba_tcls_fulfill_qa_order_create_response,omitempty"`
}

type AlibabaTclsFulfillQaOrderCreateResponse struct {

    // 鹰眼id
    TraceId   string `json:"trace_id,omitempty"`

    // 返回素材id
    Data   string `json:"data,omitempty"`

    // 是否成功
    RtSuccess   bool `json:"rt_success,omitempty"`

    // 错误码
    RtErrorCode   int64 `json:"rt_error_code,omitempty"`

    // 错误信息
    RtErrorMsg   string `json:"rt_error_msg,omitempty"`

}
