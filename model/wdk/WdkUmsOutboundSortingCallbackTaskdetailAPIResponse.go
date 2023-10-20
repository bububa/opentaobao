package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingCallbackTaskdetailAPIResponse dps分货，明细回传 API返回值
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
type WdkUmsOutboundSortingCallbackTaskdetailAPIResponse struct {
	model.CommonResponse
	WdkUmsOutboundSortingCallbackTaskdetailAPIResponseModel
}

// Reset 清空结构体
func (m *WdkUmsOutboundSortingCallbackTaskdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkUmsOutboundSortingCallbackTaskdetailAPIResponseModel).Reset()
}

// WdkUmsOutboundSortingCallbackTaskdetailAPIResponseModel is dps分货，明细回传 成功返回结果
type WdkUmsOutboundSortingCallbackTaskdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_callback_taskdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkUmsOutboundSortingCallbackTaskdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkUmsOutboundSortingCallbackTaskdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkUmsOutboundSortingCallbackTaskdetailAPIResponse)
	},
}

// GetWdkUmsOutboundSortingCallbackTaskdetailAPIResponse 从 sync.Pool 获取 WdkUmsOutboundSortingCallbackTaskdetailAPIResponse
func GetWdkUmsOutboundSortingCallbackTaskdetailAPIResponse() *WdkUmsOutboundSortingCallbackTaskdetailAPIResponse {
	return poolWdkUmsOutboundSortingCallbackTaskdetailAPIResponse.Get().(*WdkUmsOutboundSortingCallbackTaskdetailAPIResponse)
}

// ReleaseWdkUmsOutboundSortingCallbackTaskdetailAPIResponse 将 WdkUmsOutboundSortingCallbackTaskdetailAPIResponse 保存到 sync.Pool
func ReleaseWdkUmsOutboundSortingCallbackTaskdetailAPIResponse(v *WdkUmsOutboundSortingCallbackTaskdetailAPIResponse) {
	v.Reset()
	poolWdkUmsOutboundSortingCallbackTaskdetailAPIResponse.Put(v)
}
