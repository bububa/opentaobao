package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantskuCreateAPIResponse 商家商品信息新建 API返回值
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
type AlibabaWdkItemMerchantskuCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantskuCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantskuCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemMerchantskuCreateAPIResponseModel).Reset()
}

// AlibabaWdkItemMerchantskuCreateAPIResponseModel is 商家商品信息新建 成功返回结果
type AlibabaWdkItemMerchantskuCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantskuCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantskuCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemMerchantskuCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantskuCreateAPIResponse)
	},
}

// GetAlibabaWdkItemMerchantskuCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemMerchantskuCreateAPIResponse
func GetAlibabaWdkItemMerchantskuCreateAPIResponse() *AlibabaWdkItemMerchantskuCreateAPIResponse {
	return poolAlibabaWdkItemMerchantskuCreateAPIResponse.Get().(*AlibabaWdkItemMerchantskuCreateAPIResponse)
}

// ReleaseAlibabaWdkItemMerchantskuCreateAPIResponse 将 AlibabaWdkItemMerchantskuCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemMerchantskuCreateAPIResponse(v *AlibabaWdkItemMerchantskuCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemMerchantskuCreateAPIResponse.Put(v)
}
