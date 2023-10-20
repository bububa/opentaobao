package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreverserefundAPIResponse 退款打款 API返回值
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
type AlibabawdkreverserefundAPIResponse struct {
	model.CommonResponse
	AlibabawdkreverserefundAPIResponseModel
}

// AlibabawdkreverserefundAPIResponseModel is 退款打款 成功返回结果
type AlibabawdkreverserefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabawdkreverserefundResult `json:"result,omitempty" xml:"result,omitempty"`
}
