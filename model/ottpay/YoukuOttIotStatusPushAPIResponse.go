package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttIotStatusPushAPIResponse iot设备状态变化通知接口 API返回值
// youku.ott.iot.status.push
//
// ott iot设备状态通知
type YoukuOttIotStatusPushAPIResponse struct {
	model.CommonResponse
	YoukuOttIotStatusPushAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttIotStatusPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttIotStatusPushAPIResponseModel).Reset()
}

// YoukuOttIotStatusPushAPIResponseModel is iot设备状态变化通知接口 成功返回结果
type YoukuOttIotStatusPushAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_iot_status_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 成功标识
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttIotStatusPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolYoukuOttIotStatusPushAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttIotStatusPushAPIResponse)
	},
}

// GetYoukuOttIotStatusPushAPIResponse 从 sync.Pool 获取 YoukuOttIotStatusPushAPIResponse
func GetYoukuOttIotStatusPushAPIResponse() *YoukuOttIotStatusPushAPIResponse {
	return poolYoukuOttIotStatusPushAPIResponse.Get().(*YoukuOttIotStatusPushAPIResponse)
}

// ReleaseYoukuOttIotStatusPushAPIResponse 将 YoukuOttIotStatusPushAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttIotStatusPushAPIResponse(v *YoukuOttIotStatusPushAPIResponse) {
	v.Reset()
	poolYoukuOttIotStatusPushAPIResponse.Put(v)
}
