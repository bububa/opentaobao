package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse 直连医院上传接口 API返回值
// alibaba.alihealth.medicalbase.hos.syncnew
//
// 直连医院上传接口
type AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseHosSyncnewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseHosSyncnewAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseHosSyncnewAPIResponseModel is 直连医院上传接口 成功返回结果
type AlibabaAlihealthMedicalbaseHosSyncnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_hos_syncnew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseHosSyncnewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse
func GetAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse() *AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse {
	return poolAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse.Get().(*AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse 将 AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse(v *AlibabaAlihealthMedicalbaseHosSyncnewAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseHosSyncnewAPIResponse.Put(v)
}
