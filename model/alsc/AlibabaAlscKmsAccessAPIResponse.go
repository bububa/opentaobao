package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscKmsAccessAPIResponse 本地生活风控数据接入 API返回值
// alibaba.alsc.kms.access
//
// 第三方使用本地生活数据对外提供服务，上报访问日志信息接口
type AlibabaAlscKmsAccessAPIResponse struct {
	model.CommonResponse
	AlibabaAlscKmsAccessAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscKmsAccessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscKmsAccessAPIResponseModel).Reset()
}

// AlibabaAlscKmsAccessAPIResponseModel is 本地生活风控数据接入 成功返回结果
type AlibabaAlscKmsAccessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_kms_access_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// code
	Resultcode string `json:"resultcode,omitempty" xml:"resultcode,omitempty"`
	// message
	Resultmessage string `json:"resultmessage,omitempty" xml:"resultmessage,omitempty"`
	// 是否成功
	Resultsuccess bool `json:"resultsuccess,omitempty" xml:"resultsuccess,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscKmsAccessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resultcode = ""
	m.Resultmessage = ""
	m.Resultsuccess = false
}

var poolAlibabaAlscKmsAccessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscKmsAccessAPIResponse)
	},
}

// GetAlibabaAlscKmsAccessAPIResponse 从 sync.Pool 获取 AlibabaAlscKmsAccessAPIResponse
func GetAlibabaAlscKmsAccessAPIResponse() *AlibabaAlscKmsAccessAPIResponse {
	return poolAlibabaAlscKmsAccessAPIResponse.Get().(*AlibabaAlscKmsAccessAPIResponse)
}

// ReleaseAlibabaAlscKmsAccessAPIResponse 将 AlibabaAlscKmsAccessAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscKmsAccessAPIResponse(v *AlibabaAlscKmsAccessAPIResponse) {
	v.Reset()
	poolAlibabaAlscKmsAccessAPIResponse.Put(v)
}
