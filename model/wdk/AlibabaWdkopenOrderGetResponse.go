package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商户订单获取 APIResponse
alibaba.wdkopen.order.get

商户通过五道口订单id获取订单信息
*/
type AlibabaWdkopenOrderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdkopen_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果对象
    
    TopBaseResult   *TopBaseResult `json:"top_base_result,omitempty" xml:"