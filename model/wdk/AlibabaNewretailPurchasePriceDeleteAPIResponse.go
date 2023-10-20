package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNewretailPurchasePriceDeleteAPIResponse 共享库存 商户删除采购价 API返回值
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
type AlibabaNewretailPurchasePriceDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaNewretailPurchasePriceDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNewretailPurchasePriceDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNewretailPurchasePriceDeleteAPIResponseModel).Reset()
}

// AlibabaNewretailPurchasePriceDeleteAPIResponseModel is 共享库存 商户删除采购价 成功返回结果
type AlibabaNewretailPurchasePriceDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_newretail_purchase_price_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 拆单结果对象
	Result *TopBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNewretailPurchasePriceDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNewretailPurchasePriceDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNewretailPurchasePriceDeleteAPIResponse)
	},
}

// GetAlibabaNewretailPurchasePriceDeleteAPIResponse 从 sync.Pool 获取 AlibabaNewretailPurchasePriceDeleteAPIResponse
func GetAlibabaNewretailPurchasePriceDeleteAPIResponse() *AlibabaNewretailPurchasePriceDeleteAPIResponse {
	return poolAlibabaNewretailPurchasePriceDeleteAPIResponse.Get().(*AlibabaNewretailPurchasePriceDeleteAPIResponse)
}

// ReleaseAlibabaNewretailPurchasePriceDeleteAPIResponse 将 AlibabaNewretailPurchasePriceDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNewretailPurchasePriceDeleteAPIResponse(v *AlibabaNewretailPurchasePriceDeleteAPIResponse) {
	v.Reset()
	poolAlibabaNewretailPurchasePriceDeleteAPIResponse.Put(v)
}
