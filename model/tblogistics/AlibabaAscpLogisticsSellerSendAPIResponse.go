package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerSendAPIResponse 商家配送发货 API返回值
// alibaba.ascp.logistics.seller.send
//
// 该API提供商家配送发货能力，使用该接口发货，交易订单状态会直接变成卖家已发货
type AlibabaAscpLogisticsSellerSendAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsSellerSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsSellerSendAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsSellerSendAPIResponseModel is 商家配送发货 成功返回结果
type AlibabaAscpLogisticsSellerSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsSellerSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsSellerSendAPIResponse)
	},
}

// GetAlibabaAscpLogisticsSellerSendAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsSellerSendAPIResponse
func GetAlibabaAscpLogisticsSellerSendAPIResponse() *AlibabaAscpLogisticsSellerSendAPIResponse {
	return poolAlibabaAscpLogisticsSellerSendAPIResponse.Get().(*AlibabaAscpLogisticsSellerSendAPIResponse)
}

// ReleaseAlibabaAscpLogisticsSellerSendAPIResponse 将 AlibabaAscpLogisticsSellerSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsSellerSendAPIResponse(v *AlibabaAscpLogisticsSellerSendAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsSellerSendAPIResponse.Put(v)
}
