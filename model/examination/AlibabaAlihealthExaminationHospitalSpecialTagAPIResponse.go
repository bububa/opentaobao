package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse 体检机构获取特色服务标签 API返回值
// alibaba.alihealth.examination.hospital.special.tag
//
// 体检机构获取特色服务标签列表
type AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationHospitalSpecialTagAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationHospitalSpecialTagAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationHospitalSpecialTagAPIResponseModel is 体检机构获取特色服务标签 成功返回结果
type AlibabaAlihealthExaminationHospitalSpecialTagAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_hospital_special_tag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationHospitalSpecialTagAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse
func GetAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse() *AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse {
	return poolAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse.Get().(*AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse 将 AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse(v *AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationHospitalSpecialTagAPIResponse.Put(v)
}
