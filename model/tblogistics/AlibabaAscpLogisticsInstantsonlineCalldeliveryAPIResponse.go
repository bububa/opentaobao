package tblogistics

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel is 同城配送在线下单正式下单呼叫运力 成功返回结果
type AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_calldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse
func GetAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse() *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse {
	return poolAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse.Get().(*AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse 将 AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse(v *AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse.Put(v)
}
