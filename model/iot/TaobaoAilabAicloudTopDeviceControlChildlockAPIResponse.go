package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse 设备儿童锁 API返回值
// taobao.ailab.aicloud.top.device.control.childlock
//
// 设备儿童锁
type TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlChildlockAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlChildlockAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlChildlockAPIResponseModel is 设备儿童锁 成功返回结果
type TaobaoAilabAicloudTopDeviceControlChildlockAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_childlock_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 业务请求是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 网络请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlChildlockAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse() *TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse 将 TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse(v *TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlChildlockAPIResponse.Put(v)
}
