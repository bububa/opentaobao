package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCreativeRpthourlistAPIResponse 超级推荐【商品推广】创意分时报表查询 API返回值
// taobao.feedflow.item.creative.rpthourlist
//
// 创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
type TaobaoFeedflowItemCreativeRpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCreativeRpthourlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCreativeRpthourlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCreativeRpthourlistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCreativeRpthourlistAPIResponseModel is 超级推荐【商品推广】创意分时报表查询 成功返回结果
type TaobaoFeedflowItemCreativeRpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_creative_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowItemCreativeRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCreativeRpthourlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCreativeRpthourlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCreativeRpthourlistAPIResponse)
	},
}

// GetTaobaoFeedflowItemCreativeRpthourlistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCreativeRpthourlistAPIResponse
func GetTaobaoFeedflowItemCreativeRpthourlistAPIResponse() *TaobaoFeedflowItemCreativeRpthourlistAPIResponse {
	return poolTaobaoFeedflowItemCreativeRpthourlistAPIResponse.Get().(*TaobaoFeedflowItemCreativeRpthourlistAPIResponse)
}

// ReleaseTaobaoFeedflowItemCreativeRpthourlistAPIResponse 将 TaobaoFeedflowItemCreativeRpthourlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCreativeRpthourlistAPIResponse(v *TaobaoFeedflowItemCreativeRpthourlistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCreativeRpthourlistAPIResponse.Put(v)
}
