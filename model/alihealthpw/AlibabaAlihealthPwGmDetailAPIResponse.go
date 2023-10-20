package alihealthpw

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpwgmdetailAPIResponse 同情用药申请单详情接口 API返回值
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
type AlibabaalihealthpwgmdetailAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthpwgmdetailAPIResponseModel
}

// AlibabaalihealthpwgmdetailAPIResponseModel is 同情用药申请单详情接口 成功返回结果
type AlibabaalihealthpwgmdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_gm_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}
