package aesolution

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionorderinfogetAPIResponse get order detail info API返回值
// aliexpress.solution.order.info.get
//
// get order detail info
type AliexpresssolutionorderinfogetAPIResponse struct {
	model.CommonResponse
	AliexpresssolutionorderinfogetAPIResponseModel
}

// AliexpresssolutionorderinfogetAPIResponseModel is get order detail info 成功返回结果
type AliexpresssolutionorderinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
