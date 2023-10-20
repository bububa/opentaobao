package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdzoneListAPIResponse 批量查询可用广告位列表 API返回值
// taobao.feedflow.item.adzone.list
//
// 批量查询可用广告位列表
type TaobaoFeedflowItemAdzoneListAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdzoneListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdzoneListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdzoneListAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdzoneListAPIResponseModel is 批量查询可用广告位列表 成功返回结果
type TaobaoFeedflowItemAdzoneListAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adzone_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdzoneListResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdzoneListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdzoneListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdzoneListAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdzoneListAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdzoneListAPIResponse
func GetTaobaoFeedflowItemAdzoneListAPIResponse() *TaobaoFeedflowItemAdzoneListAPIResponse {
	return poolTaobaoFeedflowItemAdzoneListAPIResponse.Get().(*TaobaoFeedflowItemAdzoneListAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdzoneListAPIResponse 将 TaobaoFeedflowItemAdzoneListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdzoneListAPIResponse(v *TaobaoFeedflowItemAdzoneListAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdzoneListAPIResponse.Put(v)
}
