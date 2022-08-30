package wdk

import (
	"encoding/xml"

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

// WdkUmsOutboundSortingCanclearareaAPIResponseModel is dps分货-是否能够清场 成功返回结果
type WdkUmsOutboundSortingCanclearareaAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_cancleararea_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
