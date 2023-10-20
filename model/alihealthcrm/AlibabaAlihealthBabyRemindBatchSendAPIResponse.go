package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBabyRemindBatchSendAPIResponse 批量发送母婴提醒 API返回值
// alibaba.alihealth.baby.remind.batch.send
//
// 批量发送母婴提醒
type AlibabaAlihealthBabyRemindBatchSendAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBabyRemindBatchSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBabyRemindBatchSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBabyRemindBatchSendAPIResponseModel).Reset()
}

// AlibabaAlihealthBabyRemindBatchSendAPIResponseModel is 批量发送母婴提醒 成功返回结果
type AlibabaAlihealthBabyRemindBatchSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_baby_remind_batch_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBabyRemindBatchSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthBabyRemindBatchSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBabyRemindBatchSendAPIResponse)
	},
}

// GetAlibabaAlihealthBabyRemindBatchSendAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBabyRemindBatchSendAPIResponse
func GetAlibabaAlihealthBabyRemindBatchSendAPIResponse() *AlibabaAlihealthBabyRemindBatchSendAPIResponse {
	return poolAlibabaAlihealthBabyRemindBatchSendAPIResponse.Get().(*AlibabaAlihealthBabyRemindBatchSendAPIResponse)
}

// ReleaseAlibabaAlihealthBabyRemindBatchSendAPIResponse 将 AlibabaAlihealthBabyRemindBatchSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBabyRemindBatchSendAPIResponse(v *AlibabaAlihealthBabyRemindBatchSendAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBabyRemindBatchSendAPIResponse.Put(v)
}
