package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowAccountRptdailylistAPIResponse 获取广告主分日数据 API返回值
// taobao.feedflow.account.rptdailylist
//
// 获取广告主分日数据
type TaobaoFeedflowAccountRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowAccountRptdailylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowAccountRptdailylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowAccountRptdailylistAPIResponseModel).Reset()
}

// TaobaoFeedflowAccountRptdailylistAPIResponseModel is 获取广告主分日数据 成功返回结果
type TaobaoFeedflowAccountRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_account_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowAccountRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowAccountRptdailylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowAccountRptdailylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowAccountRptdailylistAPIResponse)
	},
}

// GetTaobaoFeedflowAccountRptdailylistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowAccountRptdailylistAPIResponse
func GetTaobaoFeedflowAccountRptdailylistAPIResponse() *TaobaoFeedflowAccountRptdailylistAPIResponse {
	return poolTaobaoFeedflowAccountRptdailylistAPIResponse.Get().(*TaobaoFeedflowAccountRptdailylistAPIResponse)
}

// ReleaseTaobaoFeedflowAccountRptdailylistAPIResponse 将 TaobaoFeedflowAccountRptdailylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowAccountRptdailylistAPIResponse(v *TaobaoFeedflowAccountRptdailylistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowAccountRptdailylistAPIResponse.Put(v)
}
