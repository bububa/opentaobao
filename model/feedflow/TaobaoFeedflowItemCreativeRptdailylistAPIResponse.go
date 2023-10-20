package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCreativeRptdailylistAPIResponse 创意分日数据查询 API返回值
// taobao.feedflow.item.creative.rptdailylist
//
// 创意分日数据查询
type TaobaoFeedflowItemCreativeRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCreativeRptdailylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCreativeRptdailylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCreativeRptdailylistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCreativeRptdailylistAPIResponseModel is 创意分日数据查询 成功返回结果
type TaobaoFeedflowItemCreativeRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_creative_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowItemCreativeRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCreativeRptdailylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCreativeRptdailylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCreativeRptdailylistAPIResponse)
	},
}

// GetTaobaoFeedflowItemCreativeRptdailylistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCreativeRptdailylistAPIResponse
func GetTaobaoFeedflowItemCreativeRptdailylistAPIResponse() *TaobaoFeedflowItemCreativeRptdailylistAPIResponse {
	return poolTaobaoFeedflowItemCreativeRptdailylistAPIResponse.Get().(*TaobaoFeedflowItemCreativeRptdailylistAPIResponse)
}

// ReleaseTaobaoFeedflowItemCreativeRptdailylistAPIResponse 将 TaobaoFeedflowItemCreativeRptdailylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCreativeRptdailylistAPIResponse(v *TaobaoFeedflowItemCreativeRptdailylistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCreativeRptdailylistAPIResponse.Put(v)
}
