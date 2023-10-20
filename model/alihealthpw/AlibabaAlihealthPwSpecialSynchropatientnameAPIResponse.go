package alihealthpw

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse 同步患者姓名至阿里健康 API返回值
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
type AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPwSpecialSynchropatientnameAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPwSpecialSynchropatientnameAPIResponseModel).Reset()
}

// AlibabaAlihealthPwSpecialSynchropatientnameAPIResponseModel is 同步患者姓名至阿里健康 成功返回结果
type AlibabaAlihealthPwSpecialSynchropatientnameAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pw_special_synchropatientname_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResponseMessage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPwSpecialSynchropatientnameAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse)
	},
}

// GetAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse
func GetAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse() *AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse {
	return poolAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse.Get().(*AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse)
}

// ReleaseAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse 将 AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse(v *AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPwSpecialSynchropatientnameAPIResponse.Put(v)
}
