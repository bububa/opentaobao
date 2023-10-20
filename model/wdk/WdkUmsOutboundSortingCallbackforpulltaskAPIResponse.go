package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkumsoutboundsortingcallbackforpulltaskAPIResponse dps分货-任务拉取确定接口 API返回值
// wdk.ums.outbound.sorting.callbackforpulltask
//
// dps分货-任务拉取确定接口
type WdkumsoutboundsortingcallbackforpulltaskAPIResponse struct {
	model.CommonResponse
	WdkumsoutboundsortingcallbackforpulltaskAPIResponseModel
}

// WdkumsoutboundsortingcallbackforpulltaskAPIResponseModel is dps分货-任务拉取确定接口 成功返回结果
type WdkumsoutboundsortingcallbackforpulltaskAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_callbackforpulltask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
