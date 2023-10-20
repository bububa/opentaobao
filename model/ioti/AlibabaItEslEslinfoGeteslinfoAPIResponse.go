package ioti

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitesleslinfogeteslinfoAPIResponse 厂测查询价签当前信息 API返回值
// alibaba.it.esl.eslinfo.geteslinfo
//
// 工厂对生产出的电子价签进行全流程功能测试，查询价签当前上报的信息
type AlibabaitesleslinfogeteslinfoAPIResponse struct {
	model.CommonResponse
	AlibabaitesleslinfogeteslinfoAPIResponseModel
}

// AlibabaitesleslinfogeteslinfoAPIResponseModel is 厂测查询价签当前信息 成功返回结果
type AlibabaitesleslinfogeteslinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_esl_eslinfo_geteslinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	EslErrorCode string `json:"esl_error_code,omitempty" xml:"esl_error_code,omitempty"`
	// 错误信息
	EslErrorMsg string `json:"esl_error_msg,omitempty" xml:"esl_error_msg,omitempty"`
	// 数据
	Content *EslTopEngineAssetsDo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功标识
	EslSuccess bool `json:"esl_success,omitempty" xml:"esl_success,omitempty"`
}
