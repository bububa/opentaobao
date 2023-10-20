package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse 点播url API返回值
// taobao.ailab.aicloud.top.device.control.playurl
//
// 点播url
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponseModel is 点播url 成功返回结果
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_playurl_response"`
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
func (m *TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse() *TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse 将 TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse(v *TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse.Put(v)
}
