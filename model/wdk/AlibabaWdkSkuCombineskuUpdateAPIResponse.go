package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuUpdateAPIResponse 组合商品更新接口 API返回值
// alibaba.wdk.sku.combinesku.update
//
// 组合商品修改接口
type AlibabaWdkSkuCombineskuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCombineskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCombineskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCombineskuUpdateAPIResponseModel).Reset()
}

// AlibabaWdkSkuCombineskuUpdateAPIResponseModel is 组合商品更新接口 成功返回结果
type AlibabaWdkSkuCombineskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_combinesku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCombineskuUpdateApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCombineskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCombineskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCombineskuUpdateAPIResponse)
	},
}

// GetAlibabaWdkSkuCombineskuUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCombineskuUpdateAPIResponse
func GetAlibabaWdkSkuCombineskuUpdateAPIResponse() *AlibabaWdkSkuCombineskuUpdateAPIResponse {
	return poolAlibabaWdkSkuCombineskuUpdateAPIResponse.Get().(*AlibabaWdkSkuCombineskuUpdateAPIResponse)
}

// ReleaseAlibabaWdkSkuCombineskuUpdateAPIResponse 将 AlibabaWdkSkuCombineskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCombineskuUpdateAPIResponse(v *AlibabaWdkSkuCombineskuUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCombineskuUpdateAPIResponse.Put(v)
}
