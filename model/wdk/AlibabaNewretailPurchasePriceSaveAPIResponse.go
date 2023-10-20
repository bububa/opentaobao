package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNewretailPurchasePriceSaveAPIResponse 共享库存 采购价上传接口 API返回值
// alibaba.newretail.purchase.price.save
//
// 共享库存业务 供应商上传商品采购价
type AlibabaNewretailPurchasePriceSaveAPIResponse struct {
	model.CommonResponse
	AlibabaNewretailPurchasePriceSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNewretailPurchasePriceSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNewretailPurchasePriceSaveAPIResponseModel).Reset()
}

// AlibabaNewretailPurchasePriceSaveAPIResponseModel is 共享库存 采购价上传接口 成功返回结果
type AlibabaNewretailPurchasePriceSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_newretail_purchase_price_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果对象
	Result *TopBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNewretailPurchasePriceSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNewretailPurchasePriceSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNewretailPurchasePriceSaveAPIResponse)
	},
}

// GetAlibabaNewretailPurchasePriceSaveAPIResponse 从 sync.Pool 获取 AlibabaNewretailPurchasePriceSaveAPIResponse
func GetAlibabaNewretailPurchasePriceSaveAPIResponse() *AlibabaNewretailPurchasePriceSaveAPIResponse {
	return poolAlibabaNewretailPurchasePriceSaveAPIResponse.Get().(*AlibabaNewretailPurchasePriceSaveAPIResponse)
}

// ReleaseAlibabaNewretailPurchasePriceSaveAPIResponse 将 AlibabaNewretailPurchasePriceSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNewretailPurchasePriceSaveAPIResponse(v *AlibabaNewretailPurchasePriceSaveAPIResponse) {
	v.Reset()
	poolAlibabaNewretailPurchasePriceSaveAPIResponse.Put(v)
}
