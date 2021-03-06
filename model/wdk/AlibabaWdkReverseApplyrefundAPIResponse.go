package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseApplyrefundAPIResponse 逆向申请接口 API返回值
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
type AlibabaWdkReverseApplyrefundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseApplyrefundAPIResponseModel
}

// AlibabaWdkReverseApplyrefundAPIResponseModel is 逆向申请接口 成功返回结果
type AlibabaWdkReverseApplyrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_applyrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回result
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
