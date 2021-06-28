package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口商户订单获取 APIResponse
alibaba.wdkopen.order.get

商户通过五道口订单id获取订单信息
*/
type AlibabaWdkopenOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkopenOrderGetResponse `json:"alibaba_wdkopen_order_get_response,omitempty"` 
    AlibabaWdkopenOrderGetResponse
}

/* model for simplify = false
type AlibabaWdkopenOrderGetResponse struct {

    // 结果对象
    
    TopBaseResult  *struct {
        TopBaseResult  *TopBaseResult `json:"top_base_result,omitempty"`
    } `json:"top_base_result,omitempty"`
    

}
*/

type AlibabaWdkopenOrderGetResponse struct {

    // 结果对象
    TopBaseResult   *TopBaseResult `json:"top_base_result,omitempty"`

}
