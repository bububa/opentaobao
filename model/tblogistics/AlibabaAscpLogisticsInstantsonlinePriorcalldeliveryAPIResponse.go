package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIResponse 同城配送在线下单预询价 API返回值
// alibaba.ascp.logistics.instantsonline.priorcalldelivery
//
// 同城配送在线下单预询价
type AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIResponseModel
}

// AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIResponseModel is 同城配送在线下单预询价 成功返回结果
type AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_priorcalldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
