package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveIsvModifyAPIResponse ISV调TOP主动发起改期信息 API返回值
// alibaba.alihealth.examination.reserve.isv.modify
//
// 体检机构对接_ISV发起体检套餐改期
type AlibabaAlihealthExaminationReserveIsvModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveIsvModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveIsvModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveIsvModifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReserveIsvModifyAPIResponseModel is ISV调TOP主动发起改期信息 成功返回结果
type AlibabaAlihealthExaminationReserveIsvModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_isv_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveIsvModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationReserveIsvModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveIsvModifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveIsvModifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveIsvModifyAPIResponse
func GetAlibabaAlihealthExaminationReserveIsvModifyAPIResponse() *AlibabaAlihealthExaminationReserveIsvModifyAPIResponse {
	return poolAlibabaAlihealthExaminationReserveIsvModifyAPIResponse.Get().(*AlibabaAlihealthExaminationReserveIsvModifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveIsvModifyAPIResponse 将 AlibabaAlihealthExaminationReserveIsvModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveIsvModifyAPIResponse(v *AlibabaAlihealthExaminationReserveIsvModifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveIsvModifyAPIResponse.Put(v)
}
