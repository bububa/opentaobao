package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse 零配方案上报设备 API返回值
// taobao.ailab.aicloud.smarthome.top.genielink.reportdevice
//
// 零配方案中设备联网成功之后上报设备
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponseModel).Reset()
}

// TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponseModel is 零配方案上报设备 成功返回结果
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_smarthome_top_genielink_reportdevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse)
	},
}

// GetTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse
func GetTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse() *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse {
	return poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse.Get().(*TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse)
}

// ReleaseTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse 将 TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse(v *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse.Put(v)
}
