package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdzoneRptdailylistAPIResponse 资源包分日数据查询 API返回值
// taobao.feedflow.item.adzone.rptdailylist
//
// 资源包分日数据查询
type TaobaoFeedflowItemAdzoneRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdzoneRptdailylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdzoneRptdailylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdzoneRptdailylistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdzoneRptdailylistAPIResponseModel is 资源包分日数据查询 成功返回结果
type TaobaoFeedflowItemAdzoneRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adzone_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowItemAdzoneRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdzoneRptdailylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdzoneRptdailylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdzoneRptdailylistAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdzoneRptdailylistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdzoneRptdailylistAPIResponse
func GetTaobaoFeedflowItemAdzoneRptdailylistAPIResponse() *TaobaoFeedflowItemAdzoneRptdailylistAPIResponse {
	return poolTaobaoFeedflowItemAdzoneRptdailylistAPIResponse.Get().(*TaobaoFeedflowItemAdzoneRptdailylistAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdzoneRptdailylistAPIResponse 将 TaobaoFeedflowItemAdzoneRptdailylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdzoneRptdailylistAPIResponse(v *TaobaoFeedflowItemAdzoneRptdailylistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdzoneRptdailylistAPIResponse.Put(v)
}
