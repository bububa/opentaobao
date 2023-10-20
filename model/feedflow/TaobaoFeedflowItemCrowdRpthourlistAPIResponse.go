package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdRpthourlistAPIResponse 超级推荐【商品推广】定向分时报表查询 API返回值
// taobao.feedflow.item.crowd.rpthourlist
//
// 广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
type TaobaoFeedflowItemCrowdRpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdRpthourlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdRpthourlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCrowdRpthourlistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCrowdRpthourlistAPIResponseModel is 超级推荐【商品推广】定向分时报表查询 成功返回结果
type TaobaoFeedflowItemCrowdRpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowItemCrowdRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdRpthourlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCrowdRpthourlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdRpthourlistAPIResponse)
	},
}

// GetTaobaoFeedflowItemCrowdRpthourlistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCrowdRpthourlistAPIResponse
func GetTaobaoFeedflowItemCrowdRpthourlistAPIResponse() *TaobaoFeedflowItemCrowdRpthourlistAPIResponse {
	return poolTaobaoFeedflowItemCrowdRpthourlistAPIResponse.Get().(*TaobaoFeedflowItemCrowdRpthourlistAPIResponse)
}

// ReleaseTaobaoFeedflowItemCrowdRpthourlistAPIResponse 将 TaobaoFeedflowItemCrowdRpthourlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdRpthourlistAPIResponse(v *TaobaoFeedflowItemCrowdRpthourlistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdRpthourlistAPIResponse.Put(v)
}
