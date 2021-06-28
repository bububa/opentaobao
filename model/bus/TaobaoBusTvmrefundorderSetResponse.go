package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机逆向退款接口 APIResponse
taobao.bus.tvmrefundorder.set

汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。）
*/
type TaobaoBusTvmrefundorderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmrefundorder_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"