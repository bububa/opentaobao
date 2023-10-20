package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleconsignmentiiorderperformAPIResponse 寄卖V2订单履约 API返回值
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
type AlibabaidleconsignmentiiorderperformAPIResponse struct {
	model.CommonResponse
	AlibabaidleconsignmentiiorderperformAPIResponseModel
}

// AlibabaidleconsignmentiiorderperformAPIResponseModel is 寄卖V2订单履约 成功返回结果
type AlibabaidleconsignmentiiorderperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignmentii_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaidleconsignmentiiorderperformResult `json:"result,omitempty" xml:"result,omitempty"`
}
