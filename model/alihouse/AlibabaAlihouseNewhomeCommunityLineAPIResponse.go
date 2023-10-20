package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecommunitylineAPIResponse 小区上下架 API返回值
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
type AlibabaalihousenewhomecommunitylineAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomecommunitylineAPIResponseModel
}

// AlibabaalihousenewhomecommunitylineAPIResponseModel is 小区上下架 成功返回结果
type AlibabaalihousenewhomecommunitylineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_community_line_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomecommunitylineResult `json:"result,omitempty" xml:"result,omitempty"`
}
