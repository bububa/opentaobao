package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponse 同城配送在线下单正式下单呼叫运力 API返回值
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
type AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponseModel
}

// AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponseModel is 同城配送在线下单正式下单呼叫运力 成功返回结果
type AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_calldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaascplogisticsinstantsonlinecalldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
