package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantskuUpdateAPIResponse 商家商品修改 API返回值
// alibaba.wdk.item.merchantsku.update
//
// 商家商品修改
type AlibabaWdkItemMerchantskuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemMerchantskuUpdateAPIResponseModel).Reset()
}

// AlibabaWdkItemMerchantskuUpdateAPIResponseModel is 商家商品修改 成功返回结果
type AlibabaWdkItemMerchantskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantskuUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemMerchantskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantskuUpdateAPIResponse)
	},
}

// GetAlibabaWdkItemMerchantskuUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemMerchantskuUpdateAPIResponse
func GetAlibabaWdkItemMerchantskuUpdateAPIResponse() *AlibabaWdkItemMerchantskuUpdateAPIResponse {
	return poolAlibabaWdkItemMerchantskuUpdateAPIResponse.Get().(*AlibabaWdkItemMerchantskuUpdateAPIResponse)
}

// ReleaseAlibabaWdkItemMerchantskuUpdateAPIResponse 将 AlibabaWdkItemMerchantskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemMerchantskuUpdateAPIResponse(v *AlibabaWdkItemMerchantskuUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemMerchantskuUpdateAPIResponse.Put(v)
}
