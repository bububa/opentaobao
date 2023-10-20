package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkumsoutboundsortingcallbacktaskdetailAPIResponse dps分货，明细回传 API返回值
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
type WdkumsoutboundsortingcallbacktaskdetailAPIResponse struct {
	model.CommonResponse
	WdkumsoutboundsortingcallbacktaskdetailAPIResponseModel
}

// WdkumsoutboundsortingcallbacktaskdetailAPIResponseModel is dps分货，明细回传 成功返回结果
type WdkumsoutboundsortingcallbacktaskdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_callback_taskdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
