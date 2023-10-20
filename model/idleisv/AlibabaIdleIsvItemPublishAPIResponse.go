package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemPublishAPIResponse 服务商闲鱼商品发布 API返回值
// alibaba.idle.isv.item.publish
//
// 服务商ISV闲鱼商品发布
type AlibabaIdleIsvItemPublishAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvItemPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvItemPublishAPIResponseModel).Reset()
}

// AlibabaIdleIsvItemPublishAPIResponseModel is 服务商闲鱼商品发布 成功返回结果
type AlibabaIdleIsvItemPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *IdleResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvItemPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemPublishAPIResponse)
	},
}

// GetAlibabaIdleIsvItemPublishAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvItemPublishAPIResponse
func GetAlibabaIdleIsvItemPublishAPIResponse() *AlibabaIdleIsvItemPublishAPIResponse {
	return poolAlibabaIdleIsvItemPublishAPIResponse.Get().(*AlibabaIdleIsvItemPublishAPIResponse)
}

// ReleaseAlibabaIdleIsvItemPublishAPIResponse 将 AlibabaIdleIsvItemPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvItemPublishAPIResponse(v *AlibabaIdleIsvItemPublishAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvItemPublishAPIResponse.Put(v)
}
