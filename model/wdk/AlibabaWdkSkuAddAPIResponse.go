package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuAddAPIResponse 新增商品 API返回值
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
type AlibabaWdkSkuAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuAddAPIResponseModel).Reset()
}

// AlibabaWdkSkuAddAPIResponseModel is 新增商品 成功返回结果
type AlibabaWdkSkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuAddApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuAddAPIResponse)
	},
}

// GetAlibabaWdkSkuAddAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuAddAPIResponse
func GetAlibabaWdkSkuAddAPIResponse() *AlibabaWdkSkuAddAPIResponse {
	return poolAlibabaWdkSkuAddAPIResponse.Get().(*AlibabaWdkSkuAddAPIResponse)
}

// ReleaseAlibabaWdkSkuAddAPIResponse 将 AlibabaWdkSkuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuAddAPIResponse(v *AlibabaWdkSkuAddAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuAddAPIResponse.Put(v)
}
