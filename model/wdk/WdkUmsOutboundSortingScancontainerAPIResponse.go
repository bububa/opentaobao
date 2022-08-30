package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingScancontainerAPIResponse dps分货-扫描分货容器判断是否可用 API返回值
// wdk.ums.outbound.sorting.scancontainer
//
// dps分货-扫描分货容器判断是否可用
type WdkUmsOutboundSortingScancontainerAPIResponse struct {
	model.CommonResponse
	WdkUmsOutboundSortingScancontainerAPIResponseModel
}

// WdkUmsOutboundSortingScancontainerAPIResponseModel is dps分货-扫描分货容器判断是否可用 成功返回结果
type WdkUmsOutboundSortingScancontainerAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_scancontainer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
