package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse 同城配送在线下单取消下单取消呼叫的运力 API返回值
// alibaba.ascp.logistics.instantsonline.canceldelivery
//
// 同城配送在线下单取消下单取消呼叫的运力
type AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponseModel is 同城配送在线下单取消下单取消呼叫的运力 成功返回结果
type AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_instantsonline_canceldelivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse
func GetAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse() *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse {
	return poolAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse.Get().(*AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse 将 AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse(v *AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse.Put(v)
}
