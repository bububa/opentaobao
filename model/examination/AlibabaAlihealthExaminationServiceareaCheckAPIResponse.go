package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationServiceareaCheckAPIResponse 体检机构对接_上门检测服务范围查询 API返回值
// alibaba.alihealth.examination.servicearea.check
//
// 体检机构对接_上门检测服务范围查询
type AlibabaAlihealthExaminationServiceareaCheckAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationServiceareaCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationServiceareaCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationServiceareaCheckAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationServiceareaCheckAPIResponseModel is 体检机构对接_上门检测服务范围查询 成功返回结果
type AlibabaAlihealthExaminationServiceareaCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_servicearea_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 校验结果：校验成功、校验失败
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 校验结果编码，校验成功200、校验失败400
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationServiceareaCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResponseCode = ""
}

var poolAlibabaAlihealthExaminationServiceareaCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationServiceareaCheckAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationServiceareaCheckAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationServiceareaCheckAPIResponse
func GetAlibabaAlihealthExaminationServiceareaCheckAPIResponse() *AlibabaAlihealthExaminationServiceareaCheckAPIResponse {
	return poolAlibabaAlihealthExaminationServiceareaCheckAPIResponse.Get().(*AlibabaAlihealthExaminationServiceareaCheckAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationServiceareaCheckAPIResponse 将 AlibabaAlihealthExaminationServiceareaCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationServiceareaCheckAPIResponse(v *AlibabaAlihealthExaminationServiceareaCheckAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationServiceareaCheckAPIResponse.Put(v)
}
