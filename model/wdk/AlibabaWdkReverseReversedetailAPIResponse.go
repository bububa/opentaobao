package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreversereversedetailAPIResponse 退款详情 API返回值
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
type AlibabawdkreversereversedetailAPIResponse struct {
	model.CommonResponse
	AlibabawdkreversereversedetailAPIResponseModel
}

// AlibabawdkreversereversedetailAPIResponseModel is 退款详情 成功返回结果
type AlibabawdkreversereversedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_reversedetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
