package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTagItemsQueryAPIResponse 打标结果查询-标维度 API返回值
// taobao.qimen.tag.items.query
//
// 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoQimenTagItemsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTagItemsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenTagItemsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenTagItemsQueryAPIResponseModel).Reset()
}

// TaobaoQimenTagItemsQueryAPIResponseModel is 打标结果查询-标维度 成功返回结果
type TaobaoQimenTagItemsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_tag_items_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// itemIds
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// flag
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// tagType
	TagType string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenTagItemsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemIds = m.ItemIds[:0]
	m.Flag = ""
	m.Message = ""
	m.TagType = ""
}

var poolTaobaoQimenTagItemsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTagItemsQueryAPIResponse)
	},
}

// GetTaobaoQimenTagItemsQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenTagItemsQueryAPIResponse
func GetTaobaoQimenTagItemsQueryAPIResponse() *TaobaoQimenTagItemsQueryAPIResponse {
	return poolTaobaoQimenTagItemsQueryAPIResponse.Get().(*TaobaoQimenTagItemsQueryAPIResponse)
}

// ReleaseTaobaoQimenTagItemsQueryAPIResponse 将 TaobaoQimenTagItemsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenTagItemsQueryAPIResponse(v *TaobaoQimenTagItemsQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenTagItemsQueryAPIResponse.Put(v)
}
