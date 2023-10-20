package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlCustomAPIResponse 设备控制自定义扩展接口 API返回值
// taobao.ailab.aicloud.top.device.control.custom
//
// 设备控制自定义扩展接口
type TaobaoAilabAicloudTopDeviceControlCustomAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlCustomAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlCustomAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlCustomAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlCustomAPIResponseModel is 设备控制自定义扩展接口 成功返回结果
type TaobaoAilabAicloudTopDeviceControlCustomAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_custom_response"`
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
func (m *TaobaoAilabAicloudTopDeviceControlCustomAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceControlCustomAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlCustomAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlCustomAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlCustomAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlCustomAPIResponse() *TaobaoAilabAicloudTopDeviceControlCustomAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlCustomAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlCustomAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlCustomAPIResponse 将 TaobaoAilabAicloudTopDeviceControlCustomAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlCustomAPIResponse(v *TaobaoAilabAicloudTopDeviceControlCustomAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlCustomAPIResponse.Put(v)
}
