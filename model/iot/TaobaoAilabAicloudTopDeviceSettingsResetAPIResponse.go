package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse 重置设备个性化设置 API返回值
// taobao.ailab.aicloud.top.device.settings.reset
//
// 重置设备个性化设置
type TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceSettingsResetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceSettingsResetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceSettingsResetAPIResponseModel is 重置设备个性化设置 成功返回结果
type TaobaoAilabAicloudTopDeviceSettingsResetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_settings_reset_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 扩展字段
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 业务结果是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 网络请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceSettingsResetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ExtendInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse
func GetTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse() *TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse 将 TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse(v *TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceSettingsResetAPIResponse.Put(v)
}
