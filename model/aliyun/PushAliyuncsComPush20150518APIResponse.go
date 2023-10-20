package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// PushAliyuncsComPush20150518APIResponse 云推送指令API API返回值
// push.aliyuncs.com.push.20150518
//
// 阿里云推送新增API，允许一条推送指令同时发布到多个终端上。
type PushAliyuncsComPush20150518APIResponse struct {
	model.CommonResponse
	PushAliyuncsComPush20150518APIResponseModel
}

// Reset 清空结构体
func (m *PushAliyuncsComPush20150518APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.PushAliyuncsComPush20150518APIResponseModel).Reset()
}

// PushAliyuncsComPush20150518APIResponseModel is 云推送指令API 成功返回结果
type PushAliyuncsComPush20150518APIResponseModel struct {
	XMLName xml.Name `xml:"push_aliyuncs_com_push_20150518_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息ID,用于查询
	ResponseParams string `json:"responseParams,omitempty" xml:"responseParams,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *PushAliyuncsComPush20150518APIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseParams = ""
	m.Success = false
}

var poolPushAliyuncsComPush20150518APIResponse = sync.Pool{
	New: func() any {
		return new(PushAliyuncsComPush20150518APIResponse)
	},
}

// GetPushAliyuncsComPush20150518APIResponse 从 sync.Pool 获取 PushAliyuncsComPush20150518APIResponse
func GetPushAliyuncsComPush20150518APIResponse() *PushAliyuncsComPush20150518APIResponse {
	return poolPushAliyuncsComPush20150518APIResponse.Get().(*PushAliyuncsComPush20150518APIResponse)
}

// ReleasePushAliyuncsComPush20150518APIResponse 将 PushAliyuncsComPush20150518APIResponse 保存到 sync.Pool
func ReleasePushAliyuncsComPush20150518APIResponse(v *PushAliyuncsComPush20150518APIResponse) {
	v.Reset()
	poolPushAliyuncsComPush20150518APIResponse.Put(v)
}
