package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse 通知改期结果 API返回值
// alibaba.alihealth.examination.reserve.modify.notify
//
// 体检状态为改期中，服务上通知健康是否改期成功
type AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveModifyNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveModifyNotifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReserveModifyNotifyAPIResponseModel is 通知改期结果 成功返回结果
type AlibabaAlihealthExaminationReserveModifyNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_modify_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveModifyNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse
func GetAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse() *AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse {
	return poolAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse.Get().(*AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse 将 AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse(v *AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveModifyNotifyAPIResponse.Put(v)
}
