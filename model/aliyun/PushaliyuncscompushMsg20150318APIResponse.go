package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPushMsg20150318APIResponse 消息推送 API返回值
// push.aliyuncs.com.pushMsg.2015-03-18
//
// 消息推送  ,支持指定用户/账号/广播等模式
type PushAliyuncsComPushMsg20150318APIResponse struct {
	model.CommonResponse
	PushAliyuncsComPushMsg20150318APIResponseModel
}

// Reset 清空结构体
func (m *PushAliyuncsComPushMsg20150318APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.PushAliyuncsComPushMsg20150318APIResponseModel).Reset()
}

// PushAliyuncsComPushMsg20150318APIResponseModel is 消息推送 成功返回结果
type PushAliyuncsComPushMsg20150318APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_pushMsg_2015-03-18_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息ID
	ResponseParams int64 `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *PushAliyuncsComPushMsg20150318APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseParams = 0
	m.Success = false
}

var poolPushAliyuncsComPushMsg20150318APIResponse = sync.Pool{
	New: func() any {
		return new(PushAliyuncsComPushMsg20150318APIResponse)
	},
}

// GetPushAliyuncsComPushMsg20150318APIResponse 从 sync.Pool 获取 PushAliyuncsComPushMsg20150318APIResponse
func GetPushAliyuncsComPushMsg20150318APIResponse() *PushAliyuncsComPushMsg20150318APIResponse {
	return poolPushAliyuncsComPushMsg20150318APIResponse.Get().(*PushAliyuncsComPushMsg20150318APIResponse)
}

// ReleasePushAliyuncsComPushMsg20150318APIResponse 将 PushAliyuncsComPushMsg20150318APIResponse 保存到 sync.Pool
func ReleasePushAliyuncsComPushMsg20150318APIResponse(v *PushAliyuncsComPushMsg20150318APIResponse) {
	v.Reset()
	poolPushAliyuncsComPushMsg20150318APIResponse.Put(v)
}
