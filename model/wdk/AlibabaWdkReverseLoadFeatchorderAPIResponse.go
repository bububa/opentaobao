package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseLoadFeatchorderAPIResponse 取货详情 API返回值
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
type AlibabaWdkReverseLoadFeatchorderAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseLoadFeatchorderAPIResponseModel
}

// AlibabaWdkReverseLoadFeatchorderAPIResponseModel is 取货详情 成功返回结果
type AlibabaWdkReverseLoadFeatchorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_load_featchorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ReverseResult
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
