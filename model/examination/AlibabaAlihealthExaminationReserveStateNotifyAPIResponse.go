package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveStateNotifyAPIResponse 体检机构对接_体检状态主动通知 API返回值
// alibaba.alihealth.examination.reserve.state.notify
//
// 到了体检当天后，服务商主动通知体检预约状态
type AlibabaAlihealthExaminationReserveStateNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveStateNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveStateNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveStateNotifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReserveStateNotifyAPIResponseModel is 体检机构对接_体检状态主动通知 成功返回结果
type AlibabaAlihealthExaminationReserveStateNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_state_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveStateNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationReserveStateNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveStateNotifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveStateNotifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveStateNotifyAPIResponse
func GetAlibabaAlihealthExaminationReserveStateNotifyAPIResponse() *AlibabaAlihealthExaminationReserveStateNotifyAPIResponse {
	return poolAlibabaAlihealthExaminationReserveStateNotifyAPIResponse.Get().(*AlibabaAlihealthExaminationReserveStateNotifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveStateNotifyAPIResponse 将 AlibabaAlihealthExaminationReserveStateNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveStateNotifyAPIResponse(v *AlibabaAlihealthExaminationReserveStateNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveStateNotifyAPIResponse.Put(v)
}
