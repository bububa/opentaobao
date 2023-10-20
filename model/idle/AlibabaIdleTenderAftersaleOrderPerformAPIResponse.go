package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderaftersaleorderperformAPIResponse 闲鱼帮卖售后订单履约 API返回值
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
type AlibabaidletenderaftersaleorderperformAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderaftersaleorderperformAPIResponseModel
}

// AlibabaidletenderaftersaleorderperformAPIResponseModel is 闲鱼帮卖售后订单履约 成功返回结果
type AlibabaidletenderaftersaleorderperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_aftersale_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	PerformError string `json:"perform_error,omitempty" xml:"perform_error,omitempty"`
	// 是否履约成功
	TransformSuccess bool `json:"transform_success,omitempty" xml:"transform_success,omitempty"`
}
