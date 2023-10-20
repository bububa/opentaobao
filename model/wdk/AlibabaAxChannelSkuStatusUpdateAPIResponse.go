package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAxChannelSkuStatusUpdateAPIResponse 翱象渠道商品上下架接口 API返回值
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
type AlibabaAxChannelSkuStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAxChannelSkuStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAxChannelSkuStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAxChannelSkuStatusUpdateAPIResponseModel).Reset()
}

// AlibabaAxChannelSkuStatusUpdateAPIResponseModel is 翱象渠道商品上下架接口 成功返回结果
type AlibabaAxChannelSkuStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ax_channel_sku_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用接口返回结果
	ApiResult *AlibabaAxChannelSkuStatusUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAxChannelSkuStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAxChannelSkuStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAxChannelSkuStatusUpdateAPIResponse)
	},
}

// GetAlibabaAxChannelSkuStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaAxChannelSkuStatusUpdateAPIResponse
func GetAlibabaAxChannelSkuStatusUpdateAPIResponse() *AlibabaAxChannelSkuStatusUpdateAPIResponse {
	return poolAlibabaAxChannelSkuStatusUpdateAPIResponse.Get().(*AlibabaAxChannelSkuStatusUpdateAPIResponse)
}

// ReleaseAlibabaAxChannelSkuStatusUpdateAPIResponse 将 AlibabaAxChannelSkuStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAxChannelSkuStatusUpdateAPIResponse(v *AlibabaAxChannelSkuStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAxChannelSkuStatusUpdateAPIResponse.Put(v)
}
