package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentOrderPerformAPIResponse
帮卖订单履约 API返回值
alibaba.idle.consignment.order.perform

帮卖订单履约，回收商同步订单信息，驱动交易流转 */
type AlibabaIdleConsignmentOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleConsignmentOrderPerformAPIResponseModel
}

// AlibabaIdleConsignmentOrderPerformAPIResponseModel is 帮卖订单履约 成功返回结果
type AlibabaIdleConsignmentOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignment_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleConsignmentOrderPerformResult `json:"result,omitempty" xml:"result,omitempty"`
}
