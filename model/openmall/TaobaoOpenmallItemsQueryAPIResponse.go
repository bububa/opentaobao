package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemsQueryAPIResponse 批量获取商品列表 API返回值
// taobao.openmall.items.query
//
// 批量获取对联盟开放的商品列表。
type TaobaoOpenmallItemsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallItemsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallItemsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallItemsQueryAPIResponseModel).Reset()
}

// TaobaoOpenmallItemsQueryAPIResponseModel is 批量获取商品列表 成功返回结果
type TaobaoOpenmallItemsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_items_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoOpenmallItemsQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallItemsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenmallItemsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallItemsQueryAPIResponse)
	},
}

// GetTaobaoOpenmallItemsQueryAPIResponse 从 sync.Pool 获取 TaobaoOpenmallItemsQueryAPIResponse
func GetTaobaoOpenmallItemsQueryAPIResponse() *TaobaoOpenmallItemsQueryAPIResponse {
	return poolTaobaoOpenmallItemsQueryAPIResponse.Get().(*TaobaoOpenmallItemsQueryAPIResponse)
}

// ReleaseTaobaoOpenmallItemsQueryAPIResponse 将 TaobaoOpenmallItemsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallItemsQueryAPIResponse(v *TaobaoOpenmallItemsQueryAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallItemsQueryAPIResponse.Put(v)
}
