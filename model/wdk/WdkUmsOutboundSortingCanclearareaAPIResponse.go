package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingCanclearareaAPIResponse dps分货-是否能够清场 API返回值
// wdk.ums.outbound.sorting.cancleararea
//
// dps分货-是否能够清场
type WdkUmsOutboundSortingCanclearareaAPIResponse struct {
	model.CommonResponse
	WdkUmsOutboundSortingCanclearareaAPIResponseModel
}

// Reset 清空结构体
func (m *WdkUmsOutboundSortingCanclearareaAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkUmsOutboundSortingCanclearareaAPIResponseModel).Reset()
}

// WdkUmsOutboundSortingCanclearareaAPIResponseModel is dps分货-是否能够清场 成功返回结果
type WdkUmsOutboundSortingCanclearareaAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_cancleararea_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkUmsOutboundSortingCanclearareaAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkUmsOutboundSortingCanclearareaAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkUmsOutboundSortingCanclearareaAPIResponse)
	},
}

// GetWdkUmsOutboundSortingCanclearareaAPIResponse 从 sync.Pool 获取 WdkUmsOutboundSortingCanclearareaAPIResponse
func GetWdkUmsOutboundSortingCanclearareaAPIResponse() *WdkUmsOutboundSortingCanclearareaAPIResponse {
	return poolWdkUmsOutboundSortingCanclearareaAPIResponse.Get().(*WdkUmsOutboundSortingCanclearareaAPIResponse)
}

// ReleaseWdkUmsOutboundSortingCanclearareaAPIResponse 将 WdkUmsOutboundSortingCanclearareaAPIResponse 保存到 sync.Pool
func ReleaseWdkUmsOutboundSortingCanclearareaAPIResponse(v *WdkUmsOutboundSortingCanclearareaAPIResponse) {
	v.Reset()
	poolWdkUmsOutboundSortingCanclearareaAPIResponse.Put(v)
}
