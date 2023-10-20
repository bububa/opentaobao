package ioti

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslEslinfoGeteslinfoAPIResponse 厂测查询价签当前信息 API返回值
// alibaba.it.esl.eslinfo.geteslinfo
//
// 工厂对生产出的电子价签进行全流程功能测试，查询价签当前上报的信息
type AlibabaItEslEslinfoGeteslinfoAPIResponse struct {
	model.CommonResponse
	AlibabaItEslEslinfoGeteslinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItEslEslinfoGeteslinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItEslEslinfoGeteslinfoAPIResponseModel).Reset()
}

// AlibabaItEslEslinfoGeteslinfoAPIResponseModel is 厂测查询价签当前信息 成功返回结果
type AlibabaItEslEslinfoGeteslinfoAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaItEslEslinfoGeteslinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EslErrorCode = ""
	m.EslErrorMsg = ""
	m.Content = nil
	m.EslSuccess = false
}

var poolAlibabaItEslEslinfoGeteslinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItEslEslinfoGeteslinfoAPIResponse)
	},
}

// GetAlibabaItEslEslinfoGeteslinfoAPIResponse 从 sync.Pool 获取 AlibabaItEslEslinfoGeteslinfoAPIResponse
func GetAlibabaItEslEslinfoGeteslinfoAPIResponse() *AlibabaItEslEslinfoGeteslinfoAPIResponse {
	return poolAlibabaItEslEslinfoGeteslinfoAPIResponse.Get().(*AlibabaItEslEslinfoGeteslinfoAPIResponse)
}

// ReleaseAlibabaItEslEslinfoGeteslinfoAPIResponse 将 AlibabaItEslEslinfoGeteslinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItEslEslinfoGeteslinfoAPIResponse(v *AlibabaItEslEslinfoGeteslinfoAPIResponse) {
	v.Reset()
	poolAlibabaItEslEslinfoGeteslinfoAPIResponse.Put(v)
}
