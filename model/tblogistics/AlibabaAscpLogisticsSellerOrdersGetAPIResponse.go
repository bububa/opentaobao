package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerOrdersGetAPIResponse 商家配送核销订单列表查询 API返回值
// alibaba.ascp.logistics.seller.orders.get
//
// 商家配送核销订单列表查询
type AlibabaAscpLogisticsSellerOrdersGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsSellerOrdersGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerOrdersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsSellerOrdersGetAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsSellerOrdersGetAPIResponseModel is 商家配送核销订单列表查询 成功返回结果
type AlibabaAscpLogisticsSellerOrdersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerOrdersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsSellerOrdersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsSellerOrdersGetAPIResponse)
	},
}

// GetAlibabaAscpLogisticsSellerOrdersGetAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsSellerOrdersGetAPIResponse
func GetAlibabaAscpLogisticsSellerOrdersGetAPIResponse() *AlibabaAscpLogisticsSellerOrdersGetAPIResponse {
	return poolAlibabaAscpLogisticsSellerOrdersGetAPIResponse.Get().(*AlibabaAscpLogisticsSellerOrdersGetAPIResponse)
}

// ReleaseAlibabaAscpLogisticsSellerOrdersGetAPIResponse 将 AlibabaAscpLogisticsSellerOrdersGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsSellerOrdersGetAPIResponse(v *AlibabaAscpLogisticsSellerOrdersGetAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsSellerOrdersGetAPIResponse.Put(v)
}
