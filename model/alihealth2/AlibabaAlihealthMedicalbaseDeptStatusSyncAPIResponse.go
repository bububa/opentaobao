package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse 挂号科室上下线 API返回值
// alibaba.alihealth.medicalbase.dept.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel
}

// AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel is 挂号科室上下线 成功返回结果
type AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_dept_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
