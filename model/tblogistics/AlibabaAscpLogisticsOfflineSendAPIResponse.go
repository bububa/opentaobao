package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsOfflineSendAPIResponse 自己联系物流发货 API返回值
// alibaba.ascp.logistics.offline.send
//
// 用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
type AlibabaAscpLogisticsOfflineSendAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsOfflineSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsOfflineSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsOfflineSendAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsOfflineSendAPIResponseModel is 自己联系物流发货 成功返回结果
type AlibabaAscpLogisticsOfflineSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_offline_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaAscpLogisticsOfflineSendResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsOfflineSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsOfflineSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsOfflineSendAPIResponse)
	},
}

// GetAlibabaAscpLogisticsOfflineSendAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsOfflineSendAPIResponse
func GetAlibabaAscpLogisticsOfflineSendAPIResponse() *AlibabaAscpLogisticsOfflineSendAPIResponse {
	return poolAlibabaAscpLogisticsOfflineSendAPIResponse.Get().(*AlibabaAscpLogisticsOfflineSendAPIResponse)
}

// ReleaseAlibabaAscpLogisticsOfflineSendAPIResponse 将 AlibabaAscpLogisticsOfflineSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsOfflineSendAPIResponse(v *AlibabaAscpLogisticsOfflineSendAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsOfflineSendAPIResponse.Put(v)
}
