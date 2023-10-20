package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecommunitysubmitAPIResponse 提交小区信息 API返回值
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
type AlibabaalihousenewhomecommunitysubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomecommunitysubmitAPIResponseModel
}

// AlibabaalihousenewhomecommunitysubmitAPIResponseModel is 提交小区信息 成功返回结果
type AlibabaalihousenewhomecommunitysubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_community_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomecommunitysubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
