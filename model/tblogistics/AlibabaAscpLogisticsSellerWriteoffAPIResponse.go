package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerWriteoffAPIResponse 商家配送核销 API返回值
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
type AlibabaAscpLogisticsSellerWriteoffAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsSellerWriteoffAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerWriteoffAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsSellerWriteoffAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsSellerWriteoffAPIResponseModel is 商家配送核销 成功返回结果
type AlibabaAscpLogisticsSellerWriteoffAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_writeoff_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerWriteoffAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsSellerWriteoffAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsSellerWriteoffAPIResponse)
	},
}

// GetAlibabaAscpLogisticsSellerWriteoffAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsSellerWriteoffAPIResponse
func GetAlibabaAscpLogisticsSellerWriteoffAPIResponse() *AlibabaAscpLogisticsSellerWriteoffAPIResponse {
	return poolAlibabaAscpLogisticsSellerWriteoffAPIResponse.Get().(*AlibabaAscpLogisticsSellerWriteoffAPIResponse)
}

// ReleaseAlibabaAscpLogisticsSellerWriteoffAPIResponse 将 AlibabaAscpLogisticsSellerWriteoffAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsSellerWriteoffAPIResponse(v *AlibabaAscpLogisticsSellerWriteoffAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsSellerWriteoffAPIResponse.Put(v)
}
