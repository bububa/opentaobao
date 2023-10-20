package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPurchasePriceAPIResponse rt回传采购价 API返回值
// alibaba.wdk.purchase.price
//
// 猫超共享库存项目-rt回传采购价
type AlibabaWdkPurchasePriceAPIResponse struct {
	model.CommonResponse
	AlibabaWdkPurchasePriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkPurchasePriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkPurchasePriceAPIResponseModel).Reset()
}

// AlibabaWdkPurchasePriceAPIResponseModel is rt回传采购价 成功返回结果
type AlibabaWdkPurchasePriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_purchase_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// SYSTEM ERROR
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// ERROR
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// true
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkPurchasePriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnMsg = ""
	m.ReturnCode = ""
	m.ReturnSuccess = false
}

var poolAlibabaWdkPurchasePriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkPurchasePriceAPIResponse)
	},
}

// GetAlibabaWdkPurchasePriceAPIResponse 从 sync.Pool 获取 AlibabaWdkPurchasePriceAPIResponse
func GetAlibabaWdkPurchasePriceAPIResponse() *AlibabaWdkPurchasePriceAPIResponse {
	return poolAlibabaWdkPurchasePriceAPIResponse.Get().(*AlibabaWdkPurchasePriceAPIResponse)
}

// ReleaseAlibabaWdkPurchasePriceAPIResponse 将 AlibabaWdkPurchasePriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkPurchasePriceAPIResponse(v *AlibabaWdkPurchasePriceAPIResponse) {
	v.Reset()
	poolAlibabaWdkPurchasePriceAPIResponse.Put(v)
}
