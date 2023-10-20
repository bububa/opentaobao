package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchrosmsAPIResponse 同步短信信息至阿里健康 API返回值
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
type AlibabaAlihealthPwSpecialSynchrosmsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwSpecialSynchrosmsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel).Reset()
}

// AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel is 同步短信信息至阿里健康 成功返回结果
type AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_special_synchrosms_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwSpecialSynchrosmsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPwSpecialSynchrosmsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwSpecialSynchrosmsAPIResponse)
	},
}

// GetAlibabaAlihealthPwSpecialSynchrosmsAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwSpecialSynchrosmsAPIResponse
func GetAlibabaAlihealthPwSpecialSynchrosmsAPIResponse() *AlibabaAlihealthPwSpecialSynchrosmsAPIResponse {
	return poolAlibabaAlihealthPwSpecialSynchrosmsAPIResponse.Get().(*AlibabaAlihealthPwSpecialSynchrosmsAPIResponse)
}

// ReleaseAlibabaAlihealthPwSpecialSynchrosmsAPIResponse 将 AlibabaAlihealthPwSpecialSynchrosmsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwSpecialSynchrosmsAPIResponse(v *AlibabaAlihealthPwSpecialSynchrosmsAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwSpecialSynchrosmsAPIResponse.Put(v)
}
