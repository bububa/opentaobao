package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemDownshelfAPIResponse 服务商闲鱼商品下架 API返回值
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
type AlibabaIdleIsvItemDownshelfAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvItemDownshelfAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemDownshelfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvItemDownshelfAPIResponseModel).Reset()
}

// AlibabaIdleIsvItemDownshelfAPIResponseModel is 服务商闲鱼商品下架 成功返回结果
type AlibabaIdleIsvItemDownshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_downshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果result
	Result *AlibabaIdleIsvItemDownshelfTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemDownshelfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvItemDownshelfAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemDownshelfAPIResponse)
	},
}

// GetAlibabaIdleIsvItemDownshelfAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvItemDownshelfAPIResponse
func GetAlibabaIdleIsvItemDownshelfAPIResponse() *AlibabaIdleIsvItemDownshelfAPIResponse {
	return poolAlibabaIdleIsvItemDownshelfAPIResponse.Get().(*AlibabaIdleIsvItemDownshelfAPIResponse)
}

// ReleaseAlibabaIdleIsvItemDownshelfAPIResponse 将 AlibabaIdleIsvItemDownshelfAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvItemDownshelfAPIResponse(v *AlibabaIdleIsvItemDownshelfAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvItemDownshelfAPIResponse.Put(v)
}
