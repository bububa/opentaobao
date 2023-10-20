package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse 直连科室上传 API返回值
// alibaba.alihealth.medicalbase.dept.syncnew
//
// 直连科室上传接口
type AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponseModel is 直连科室上传 成功返回结果
type AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_dept_syncnew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse
func GetAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse() *AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse {
	return poolAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse.Get().(*AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse 将 AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse(v *AlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseDeptSyncnewAPIResponse.Put(v)
}
