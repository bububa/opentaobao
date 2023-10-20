package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpPurchasePriceCreateAPIResponse ascp采购价写入接口 API返回值
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
type AlibabaAscpPurchasePriceCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpPurchasePriceCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpPurchasePriceCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpPurchasePriceCreateAPIResponseModel).Reset()
}

// AlibabaAscpPurchasePriceCreateAPIResponseModel is ascp采购价写入接口 成功返回结果
type AlibabaAscpPurchasePriceCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_purchase_price_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpPurchasePriceCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpPurchasePriceCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpPurchasePriceCreateAPIResponse)
	},
}

// GetAlibabaAscpPurchasePriceCreateAPIResponse 从 sync.Pool 获取 AlibabaAscpPurchasePriceCreateAPIResponse
func GetAlibabaAscpPurchasePriceCreateAPIResponse() *AlibabaAscpPurchasePriceCreateAPIResponse {
	return poolAlibabaAscpPurchasePriceCreateAPIResponse.Get().(*AlibabaAscpPurchasePriceCreateAPIResponse)
}

// ReleaseAlibabaAscpPurchasePriceCreateAPIResponse 将 AlibabaAscpPurchasePriceCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpPurchasePriceCreateAPIResponse(v *AlibabaAscpPurchasePriceCreateAPIResponse) {
	v.Reset()
	poolAlibabaAscpPurchasePriceCreateAPIResponse.Put(v)
}
