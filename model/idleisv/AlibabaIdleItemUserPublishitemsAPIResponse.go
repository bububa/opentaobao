package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleItemUserPublishitemsAPIResponse 发布的商品列表 API返回值
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
type AlibabaIdleItemUserPublishitemsAPIResponse struct {
	model.CommonResponse
	AlibabaIdleItemUserPublishitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleItemUserPublishitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleItemUserPublishitemsAPIResponseModel).Reset()
}

// AlibabaIdleItemUserPublishitemsAPIResponseModel is 发布的商品列表 成功返回结果
type AlibabaIdleItemUserPublishitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_item_user_publishitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *TopPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleItemUserPublishitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleItemUserPublishitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleItemUserPublishitemsAPIResponse)
	},
}

// GetAlibabaIdleItemUserPublishitemsAPIResponse 从 sync.Pool 获取 AlibabaIdleItemUserPublishitemsAPIResponse
func GetAlibabaIdleItemUserPublishitemsAPIResponse() *AlibabaIdleItemUserPublishitemsAPIResponse {
	return poolAlibabaIdleItemUserPublishitemsAPIResponse.Get().(*AlibabaIdleItemUserPublishitemsAPIResponse)
}

// ReleaseAlibabaIdleItemUserPublishitemsAPIResponse 将 AlibabaIdleItemUserPublishitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleItemUserPublishitemsAPIResponse(v *AlibabaIdleItemUserPublishitemsAPIResponse) {
	v.Reset()
	poolAlibabaIdleItemUserPublishitemsAPIResponse.Put(v)
}
