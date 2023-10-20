package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuUpdateAPIResponse 更新商品 API返回值
// alibaba.wdk.sku.update
//
// 开放商品更新接口
type AlibabaWdkSkuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuUpdateAPIResponseModel).Reset()
}

// AlibabaWdkSkuUpdateAPIResponseModel is 更新商品 成功返回结果
type AlibabaWdkSkuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result *AlibabaWdkSkuUpdateApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuUpdateAPIResponse)
	},
}

// GetAlibabaWdkSkuUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuUpdateAPIResponse
func GetAlibabaWdkSkuUpdateAPIResponse() *AlibabaWdkSkuUpdateAPIResponse {
	return poolAlibabaWdkSkuUpdateAPIResponse.Get().(*AlibabaWdkSkuUpdateAPIResponse)
}

// ReleaseAlibabaWdkSkuUpdateAPIResponse 将 AlibabaWdkSkuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuUpdateAPIResponse(v *AlibabaWdkSkuUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuUpdateAPIResponse.Put(v)
}
