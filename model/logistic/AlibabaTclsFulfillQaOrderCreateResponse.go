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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_fulfill_qa_order_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 鹰眼id
    
    TraceId   string `json:"trace_id,omitempty" xml:"