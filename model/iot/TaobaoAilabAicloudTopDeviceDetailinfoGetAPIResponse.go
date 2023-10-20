package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse 获取设备详细信息 API返回值
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel is 获取设备详细信息 成功返回结果
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_detailinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAilabAicloudTopDeviceDetailinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse
func GetTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse() *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse 将 TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse(v *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse.Put(v)
}
