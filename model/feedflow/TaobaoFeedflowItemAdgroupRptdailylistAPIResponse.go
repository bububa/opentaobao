package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupRptdailylistAPIResponse 推广单元分日数据查询 API返回值
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
type TaobaoFeedflowItemAdgroupRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupRptdailylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupRptdailylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupRptdailylistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupRptdailylistAPIResponseModel is 推广单元分日数据查询 成功返回结果
type TaobaoFeedflowItemAdgroupRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowItemAdgroupRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupRptdailylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupRptdailylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupRptdailylistAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupRptdailylistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupRptdailylistAPIResponse
func GetTaobaoFeedflowItemAdgroupRptdailylistAPIResponse() *TaobaoFeedflowItemAdgroupRptdailylistAPIResponse {
	return poolTaobaoFeedflowItemAdgroupRptdailylistAPIResponse.Get().(*TaobaoFeedflowItemAdgroupRptdailylistAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupRptdailylistAPIResponse 将 TaobaoFeedflowItemAdgroupRptdailylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupRptdailylistAPIResponse(v *TaobaoFeedflowItemAdgroupRptdailylistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupRptdailylistAPIResponse.Put(v)
}
