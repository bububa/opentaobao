package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创单接口 APIResponse
alibaba.tcls.fulfill.qa.order.create

根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
*/
type AlibabaTclsFulfillQaOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaTclsFulfillQaOrderCreateResponse
}

type AlibabaTclsFulfillQaOrderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_tcls_fulfill_qa_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 鹰眼id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 返回素材id
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 是否成功
    
    RtSuccess   bool `json:"rt_success,omitempty" xml:"rt_success,omitempty"`

    
    // 错误码
    
    RtErrorCode   int64 `json:"rt_error_code,omitempty" xml:"rt_error_code,omitempty"`

    
    // 错误信息
    
    RtErrorMsg   string `json:"rt_error_msg,omitempty" xml:"rt_error_msg,omitempty"`

    
}
