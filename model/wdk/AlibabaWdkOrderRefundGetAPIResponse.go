package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderRefundGetAPIResponse
五道口订单退款按ID查询 API返回值
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情 */
type AlibabaWdkOrderRefundGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderRefundGetAPIResponseModel
}

// AlibabaWdkOrderRefundGetAPIResponseModel is 五道口订单退款按ID查询 成功返回结果
type AlibabaWdkOrderRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}
