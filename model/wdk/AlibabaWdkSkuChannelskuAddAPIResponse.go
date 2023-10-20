package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuChannelskuAddAPIResponse 新增渠道商品 API返回值
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
type AlibabaWdkSkuChannelskuAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuChannelskuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuChannelskuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuChannelskuAddAPIResponseModel).Reset()
}

// AlibabaWdkSkuChannelskuAddAPIResponseModel is 新增渠道商品 成功返回结果
type AlibabaWdkSkuChannelskuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkSkuChannelskuAddApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuChannelskuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuChannelskuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuAddAPIResponse)
	},
}

// GetAlibabaWdkSkuChannelskuAddAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuChannelskuAddAPIResponse
func GetAlibabaWdkSkuChannelskuAddAPIResponse() *AlibabaWdkSkuChannelskuAddAPIResponse {
	return poolAlibabaWdkSkuChannelskuAddAPIResponse.Get().(*AlibabaWdkSkuChannelskuAddAPIResponse)
}

// ReleaseAlibabaWdkSkuChannelskuAddAPIResponse 将 AlibabaWdkSkuChannelskuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuChannelskuAddAPIResponse(v *AlibabaWdkSkuChannelskuAddAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuChannelskuAddAPIResponse.Put(v)
}
