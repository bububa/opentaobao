package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuChannelskuUpdateAPIResponse 更新渠道商品 API返回值
// alibaba.wdk.sku.channelsku.update
//
// 批量更新渠道商品，商家通过Top接入
type AlibabaWdkSkuChannelskuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuChannelskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuChannelskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuChannelskuUpdateAPIResponseModel).Reset()
}

// AlibabaWdkSkuChannelskuUpdateAPIResponseModel is 更新渠道商品 成功返回结果
type AlibabaWdkSkuChannelskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuChannelskuUpdateApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuChannelskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuChannelskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuChannelskuUpdateAPIResponse)
	},
}

// GetAlibabaWdkSkuChannelskuUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuChannelskuUpdateAPIResponse
func GetAlibabaWdkSkuChannelskuUpdateAPIResponse() *AlibabaWdkSkuChannelskuUpdateAPIResponse {
	return poolAlibabaWdkSkuChannelskuUpdateAPIResponse.Get().(*AlibabaWdkSkuChannelskuUpdateAPIResponse)
}

// ReleaseAlibabaWdkSkuChannelskuUpdateAPIResponse 将 AlibabaWdkSkuChannelskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuChannelskuUpdateAPIResponse(v *AlibabaWdkSkuChannelskuUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuChannelskuUpdateAPIResponse.Put(v)
}
