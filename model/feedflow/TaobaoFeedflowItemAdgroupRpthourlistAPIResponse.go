package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupRpthourlistAPIResponse 超级推荐【商品推广】单元分时报表查询 API返回值
// taobao.feedflow.item.adgroup.rpthourlist
//
// 广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据
type TaobaoFeedflowItemAdgroupRpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupRpthourlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupRpthourlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupRpthourlistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupRpthourlistAPIResponseModel is 超级推荐【商品推广】单元分时报表查询 成功返回结果
type TaobaoFeedflowItemAdgroupRpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowItemAdgroupRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupRpthourlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupRpthourlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupRpthourlistAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupRpthourlistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupRpthourlistAPIResponse
func GetTaobaoFeedflowItemAdgroupRpthourlistAPIResponse() *TaobaoFeedflowItemAdgroupRpthourlistAPIResponse {
	return poolTaobaoFeedflowItemAdgroupRpthourlistAPIResponse.Get().(*TaobaoFeedflowItemAdgroupRpthourlistAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupRpthourlistAPIResponse 将 TaobaoFeedflowItemAdgroupRpthourlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupRpthourlistAPIResponse(v *TaobaoFeedflowItemAdgroupRpthourlistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupRpthourlistAPIResponse.Put(v)
}
