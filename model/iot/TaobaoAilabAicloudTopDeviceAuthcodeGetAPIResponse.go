package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse 获取设备授权码 API返回值
// taobao.ailab.aicloud.top.device.authcode.get
//
// 获取设备授权码
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel is 获取设备授权码 成功返回结果
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_authcode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备授权码，后续流程中所述的auth code
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgInfo错误码信息，成功返回null
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse
func GetTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse() *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse 将 TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse(v *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse.Put(v)
}
