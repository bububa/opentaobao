package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponse 同城配送在线下单取消下单取消呼叫的运力 API返回值
// alibaba.ascp.logistics.instantsonline.canceldelivery
//
// 同城配送在线下单取消下单取消呼叫的运力
type AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponseModel
}

// AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponseModel is 同城配送在线下单取消下单取消呼叫的运力 成功返回结果
type AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_canceldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaascplogisticsinstantsonlinecanceldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
