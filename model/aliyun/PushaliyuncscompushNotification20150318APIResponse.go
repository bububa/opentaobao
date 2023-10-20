package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPushNotification20150318APIResponse 推送通知 API返回值
// push.aliyuncs.com.pushNotification.2015-03-18
//
// pushNotification
type PushAliyuncsComPushNotification20150318APIResponse struct {
	model.CommonResponse
	PushAliyuncsComPushNotification20150318APIResponseModel
}

// Reset 清空结构体
func (m *PushAliyuncsComPushNotification20150318APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.PushAliyuncsComPushNotification20150318APIResponseModel).Reset()
}

// PushAliyuncsComPushNotification20150318APIResponseModel is 推送通知 成功返回结果
type PushAliyuncsComPushNotification20150318APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_pushNotification_2015-03-18_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息ID,用于查询
	ResponseParams int64 `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *PushAliyuncsComPushNotification20150318APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseParams = 0
	m.Success = false
}

var poolPushAliyuncsComPushNotification20150318APIResponse = sync.Pool{
	New: func() any {
		return new(PushAliyuncsComPushNotification20150318APIResponse)
	},
}

// GetPushAliyuncsComPushNotification20150318APIResponse 从 sync.Pool 获取 PushAliyuncsComPushNotification20150318APIResponse
func GetPushAliyuncsComPushNotification20150318APIResponse() *PushAliyuncsComPushNotification20150318APIResponse {
	return poolPushAliyuncsComPushNotification20150318APIResponse.Get().(*PushAliyuncsComPushNotification20150318APIResponse)
}

// ReleasePushAliyuncsComPushNotification20150318APIResponse 将 PushAliyuncsComPushNotification20150318APIResponse 保存到 sync.Pool
func ReleasePushAliyuncsComPushNotification20150318APIResponse(v *PushAliyuncsComPushNotification20150318APIResponse) {
	v.Reset()
	poolPushAliyuncsComPushNotification20150318APIResponse.Put(v)
}
