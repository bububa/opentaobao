package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商户订单获取 API返回值 
alibaba.wdkopen.order.get

商户通过五道口订单id获取订单信息
*/
type AlibabaWdkopenOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkopenOrderGetAPIResponseModel
}

// 五道口商户订单获取 成功返回结果
type AlibabaWdkopenOrderGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdkopen_order_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    TopBaseResult   *TopBaseResult `json:"top_base_result,omitempty" xml:"top_base_result,omitempty"`
}
