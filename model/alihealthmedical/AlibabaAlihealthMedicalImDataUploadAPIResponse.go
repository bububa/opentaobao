package alihealthmedical

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalImDataUploadAPIResponse 三方IM图片音频消息上传 API返回值
// alibaba.alihealth.medical.im.data.upload
//
// 三方IM图片音频消息上传
type AlibabaAlihealthMedicalImDataUploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalImDataUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalImDataUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalImDataUploadAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalImDataUploadAPIResponseModel is 三方IM图片音频消息上传 成功返回结果
type AlibabaAlihealthMedicalImDataUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_im_data_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalImDataUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalImDataUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalImDataUploadAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalImDataUploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalImDataUploadAPIResponse
func GetAlibabaAlihealthMedicalImDataUploadAPIResponse() *AlibabaAlihealthMedicalImDataUploadAPIResponse {
	return poolAlibabaAlihealthMedicalImDataUploadAPIResponse.Get().(*AlibabaAlihealthMedicalImDataUploadAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalImDataUploadAPIResponse 将 AlibabaAlihealthMedicalImDataUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalImDataUploadAPIResponse(v *AlibabaAlihealthMedicalImDataUploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalImDataUploadAPIResponse.Put(v)
}
