package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceUnbindAPIResponse 解绑设备 API返回值
// taobao.ailab.aicloud.top.device.unbind
//
// 解绑设备
type TaobaoAilabAicloudTopDeviceUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel is 解绑设备 成功返回结果
type TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo错误码信息，成功返回null
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 解绑是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceUnbindAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceUnbindAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceUnbindAPIResponse
func GetTaobaoAilabAicloudTopDeviceUnbindAPIResponse() *TaobaoAilabAicloudTopDeviceUnbindAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceUnbindAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceUnbindAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceUnbindAPIResponse 将 TaobaoAilabAicloudTopDeviceUnbindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceUnbindAPIResponse(v *TaobaoAilabAicloudTopDeviceUnbindAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceUnbindAPIResponse.Put(v)
}
