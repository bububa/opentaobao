package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse 同城配送在线下单正式下单呼叫运力 API返回值
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
type AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel
}

// AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel is 同城配送在线下单正式下单呼叫运力 成功返回结果
type AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_calldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
