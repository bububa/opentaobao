package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsBuyerFreightCalculateAPIResponse 提供给买家使用的运费计算接口 API返回值
// aliexpress.logistics.buyer.freight.calculate
//
// 提供给买家使用的运费计算接口
type AliexpressLogisticsBuyerFreightCalculateAPIResponse struct {
	model.CommonResponse
	AliexpressLogisticsBuyerFreightCalculateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLogisticsBuyerFreightCalculateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLogisticsBuyerFreightCalculateAPIResponseModel).Reset()
}

// AliexpressLogisticsBuyerFreightCalculateAPIResponseModel is 提供给买家使用的运费计算接口 成功返回结果
type AliexpressLogisticsBuyerFreightCalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_logistics_buyer_freight_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AeopFreightCalculateResultListResponseForBuyer `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLogisticsBuyerFreightCalculateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressLogisticsBuyerFreightCalculateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLogisticsBuyerFreightCalculateAPIResponse)
	},
}

// GetAliexpressLogisticsBuyerFreightCalculateAPIResponse 从 sync.Pool 获取 AliexpressLogisticsBuyerFreightCalculateAPIResponse
func GetAliexpressLogisticsBuyerFreightCalculateAPIResponse() *AliexpressLogisticsBuyerFreightCalculateAPIResponse {
	return poolAliexpressLogisticsBuyerFreightCalculateAPIResponse.Get().(*AliexpressLogisticsBuyerFreightCalculateAPIResponse)
}

// ReleaseAliexpressLogisticsBuyerFreightCalculateAPIResponse 将 AliexpressLogisticsBuyerFreightCalculateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLogisticsBuyerFreightCalculateAPIResponse(v *AliexpressLogisticsBuyerFreightCalculateAPIResponse) {
	v.Reset()
	poolAliexpressLogisticsBuyerFreightCalculateAPIResponse.Put(v)
}
