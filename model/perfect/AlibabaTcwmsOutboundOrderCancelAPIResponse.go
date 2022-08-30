package perfect

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundOrderCancelAPIResponse 取消出库单 API返回值
// alibaba.tcwms.outbound.order.cancel
//
// 取消出库单
type AlibabaTcwmsOutboundOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundOrderCancelAPIResponseModel
}

// AlibabaTcwmsOutboundOrderCancelAPIResponseModel is 取消出库单 成功返回结果
type AlibabaTcwmsOutboundOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *OutboundOrderCancelResponse `json:"result,omitempty" xml:"result,omitempty"`
}
