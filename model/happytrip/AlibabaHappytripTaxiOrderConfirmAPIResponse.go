package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderConfirmAPIResponse
费用确认 API返回值
alibaba.happytrip.taxi.order.confirm

1.司机点结束计费,欢行会收到正常支付待评论 回调,确认费用无误欢行可以通过此接口确认并支付。
2.如果欢行一直不调用此接口,订单会在48小时后自动支付。 */
type AlibabaHappytripTaxiOrderConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderConfirmAPIResponseModel
}

// AlibabaHappytripTaxiOrderConfirmAPIResponseModel is 费用确认 成功返回结果
type AlibabaHappytripTaxiOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 订单确认结果对象
	Data *OrderConfirmResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
