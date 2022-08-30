package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseCreatefeatchAPIResponse 逆向取货 API返回值
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
type AlibabaWdkReverseCreatefeatchAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseCreatefeatchAPIResponseModel
}

// AlibabaWdkReverseCreatefeatchAPIResponseModel is 逆向取货 成功返回结果
type AlibabaWdkReverseCreatefeatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_createfeatch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
