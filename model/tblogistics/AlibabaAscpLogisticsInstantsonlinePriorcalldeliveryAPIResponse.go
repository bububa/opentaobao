package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponse 同城配送在线下单预询价 API返回值
// alibaba.ascp.logistics.instantsonline.priorcalldelivery
//
// 同城配送在线下单预询价
type AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponseModel
}

// AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponseModel is 同城配送在线下单预询价 成功返回结果
type AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_priorcalldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaascplogisticsinstantsonlinepriorcalldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
