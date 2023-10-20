package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdzoneRpthourlistAPIResponse 超级推荐【商品推广】资源位分时报表查询 API返回值
// taobao.feedflow.item.adzone.rpthourlist
//
// 广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
type TaobaoFeedflowItemAdzoneRpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdzoneRpthourlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdzoneRpthourlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdzoneRpthourlistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdzoneRpthourlistAPIResponseModel is 超级推荐【商品推广】资源位分时报表查询 成功返回结果
type TaobaoFeedflowItemAdzoneRpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adzone_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowItemAdzoneRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdzoneRpthourlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdzoneRpthourlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdzoneRpthourlistAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdzoneRpthourlistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdzoneRpthourlistAPIResponse
func GetTaobaoFeedflowItemAdzoneRpthourlistAPIResponse() *TaobaoFeedflowItemAdzoneRpthourlistAPIResponse {
	return poolTaobaoFeedflowItemAdzoneRpthourlistAPIResponse.Get().(*TaobaoFeedflowItemAdzoneRpthourlistAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdzoneRpthourlistAPIResponse 将 TaobaoFeedflowItemAdzoneRpthourlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdzoneRpthourlistAPIResponse(v *TaobaoFeedflowItemAdzoneRpthourlistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdzoneRpthourlistAPIResponse.Put(v)
}
