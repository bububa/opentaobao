package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytRecordinfoAPIResponse 快易通健康检查 API返回值
// alibaba.alihealth.drug.kyt.recordinfo
//
// 快易通健康检查
type AlibabaAlihealthDrugKytRecordinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytRecordinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytRecordinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytRecordinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytRecordinfoAPIResponseModel is 快易通健康检查 成功返回结果
type AlibabaAlihealthDrugKytRecordinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_recordinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 对象
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytRecordinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytRecordinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytRecordinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytRecordinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytRecordinfoAPIResponse
func GetAlibabaAlihealthDrugKytRecordinfoAPIResponse() *AlibabaAlihealthDrugKytRecordinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytRecordinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytRecordinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytRecordinfoAPIResponse 将 AlibabaAlihealthDrugKytRecordinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytRecordinfoAPIResponse(v *AlibabaAlihealthDrugKytRecordinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytRecordinfoAPIResponse.Put(v)
}
