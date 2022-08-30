package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsSortingFullContainerAPIResponse dps分货-满箱 API返回值
// wdk.ums.sorting.full.container
//
// dps分货-满箱
type WdkUmsSortingFullContainerAPIResponse struct {
	model.CommonResponse
	WdkUmsSortingFullContainerAPIResponseModel
}

// WdkUmsSortingFullContainerAPIResponseModel is dps分货-满箱 成功返回结果
type WdkUmsSortingFullContainerAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_ums_sorting_full_container_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *UmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
