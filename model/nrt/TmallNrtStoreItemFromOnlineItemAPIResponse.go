package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreItemFromOnlineItemAPIResponse 基于新模型商品id查询摊位子品id API返回值
// tmall.nrt.store.item.from.online.item
//
// 基于新模型商品id查询摊位子品id
type TmallNrtStoreItemFromOnlineItemAPIResponse struct {
	model.CommonResponse
	TmallNrtStoreItemFromOnlineItemAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtStoreItemFromOnlineItemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtStoreItemFromOnlineItemAPIResponseModel).Reset()
}

// TmallNrtStoreItemFromOnlineItemAPIResponseModel is 基于新模型商品id查询摊位子品id 成功返回结果
type TmallNrtStoreItemFromOnlineItemAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_item_from_online_item_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallNrtStoreItemFromOnlineItemResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtStoreItemFromOnlineItemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtStoreItemFromOnlineItemAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtStoreItemFromOnlineItemAPIResponse)
	},
}

// GetTmallNrtStoreItemFromOnlineItemAPIResponse 从 sync.Pool 获取 TmallNrtStoreItemFromOnlineItemAPIResponse
func GetTmallNrtStoreItemFromOnlineItemAPIResponse() *TmallNrtStoreItemFromOnlineItemAPIResponse {
	return poolTmallNrtStoreItemFromOnlineItemAPIResponse.Get().(*TmallNrtStoreItemFromOnlineItemAPIResponse)
}

// ReleaseTmallNrtStoreItemFromOnlineItemAPIResponse 将 TmallNrtStoreItemFromOnlineItemAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtStoreItemFromOnlineItemAPIResponse(v *TmallNrtStoreItemFromOnlineItemAPIResponse) {
	v.Reset()
	poolTmallNrtStoreItemFromOnlineItemAPIResponse.Put(v)
}
