package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse 获取设备状态信息 API返回值
// taobao.ailab.aicloud.top.device.statusinfo.get
//
// 获取设备状态信息
type TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponseModel is 获取设备状态信息 成功返回结果
type TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_statusinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAilabAicloudTopDeviceStatusinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse
func GetTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse() *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse 将 TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse(v *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse.Put(v)
}
