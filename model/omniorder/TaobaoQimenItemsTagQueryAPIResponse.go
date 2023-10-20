package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemsTagQueryAPIResponse 打标结果查询-商品维度 API返回值
// taobao.qimen.items.tag.query
//
// 调用该接口，查询某个/某批商品上的标
type TaobaoQimenItemsTagQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemsTagQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenItemsTagQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenItemsTagQueryAPIResponseModel).Reset()
}

// TaobaoQimenItemsTagQueryAPIResponseModel is 打标结果查询-商品维度 成功返回结果
type TaobaoQimenItemsTagQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_items_tag_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// itemTags
	ItemTags []ItemTag `json:"item_tags,omitempty" xml:"item_tags>item_tag,omitempty"`
	// flag
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenItemsTagQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemTags = m.ItemTags[:0]
	m.Flag = ""
	m.Message = ""
}

var poolTaobaoQimenItemsTagQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemsTagQueryAPIResponse)
	},
}

// GetTaobaoQimenItemsTagQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenItemsTagQueryAPIResponse
func GetTaobaoQimenItemsTagQueryAPIResponse() *TaobaoQimenItemsTagQueryAPIResponse {
	return poolTaobaoQimenItemsTagQueryAPIResponse.Get().(*TaobaoQimenItemsTagQueryAPIResponse)
}

// ReleaseTaobaoQimenItemsTagQueryAPIResponse 将 TaobaoQimenItemsTagQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenItemsTagQueryAPIResponse(v *TaobaoQimenItemsTagQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenItemsTagQueryAPIResponse.Put(v)
}
