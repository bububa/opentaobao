package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowAccountRpthourlistAPIResponse 超级推荐广告主分时报表查询 API返回值
// taobao.feedflow.account.rpthourlist
//
// 广告主分时报表查询，支持广告主查询最近90天内某一天的账户维度分时报表数据
type TaobaoFeedflowAccountRpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowAccountRpthourlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowAccountRpthourlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowAccountRpthourlistAPIResponseModel).Reset()
}

// TaobaoFeedflowAccountRpthourlistAPIResponseModel is 超级推荐广告主分时报表查询 成功返回结果
type TaobaoFeedflowAccountRpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_account_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowAccountRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowAccountRpthourlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowAccountRpthourlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowAccountRpthourlistAPIResponse)
	},
}

// GetTaobaoFeedflowAccountRpthourlistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowAccountRpthourlistAPIResponse
func GetTaobaoFeedflowAccountRpthourlistAPIResponse() *TaobaoFeedflowAccountRpthourlistAPIResponse {
	return poolTaobaoFeedflowAccountRpthourlistAPIResponse.Get().(*TaobaoFeedflowAccountRpthourlistAPIResponse)
}

// ReleaseTaobaoFeedflowAccountRpthourlistAPIResponse 将 TaobaoFeedflowAccountRpthourlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowAccountRpthourlistAPIResponse(v *TaobaoFeedflowAccountRpthourlistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowAccountRpthourlistAPIResponse.Put(v)
}
