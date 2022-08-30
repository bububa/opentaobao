package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingUserqueryAPIResponse dps-查询分货作业人员信息 API返回值
// wdk.ums.outbound.sorting.userquery
//
// dps-查询分货作业人员信息
type WdkUmsOutboundSortingUserqueryAPIResponse struct {
	model.CommonResponse
	WdkUmsOutboundSortingUserqueryAPIResponseModel
}

// WdkUmsOutboundSortingUserqueryAPIResponseModel is dps-查询分货作业人员信息 成功返回结果
type WdkUmsOutboundSortingUserqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_outbound_sorting_userquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
