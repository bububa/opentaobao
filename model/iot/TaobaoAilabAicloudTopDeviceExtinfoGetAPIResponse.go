package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse 获取设备扩展信息 API返回值
// taobao.ailab.aicloud.top.device.extinfo.get
//
// 获取设备扩展信息
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponseModel is 获取设备扩展信息 成功返回结果
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_extinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAilabAicloudTopDeviceExtinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse
func GetTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse() *TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse 将 TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse(v *TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse.Put(v)
}
