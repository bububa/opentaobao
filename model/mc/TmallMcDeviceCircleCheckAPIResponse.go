package mc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMcDeviceCircleCheckAPIResponse 云码设备圈选情况查询 API返回值
// tmall.mc.device.circle.check
//
// 云码设备圈选情况查询
type TmallMcDeviceCircleCheckAPIResponse struct {
	model.CommonResponse
	TmallMcDeviceCircleCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMcDeviceCircleCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMcDeviceCircleCheckAPIResponseModel).Reset()
}

// TmallMcDeviceCircleCheckAPIResponseModel is 云码设备圈选情况查询 成功返回结果
type TmallMcDeviceCircleCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mc_device_circle_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备相关投放计划
	Results []TaskDto `json:"results,omitempty" xml:"results>task_dto,omitempty"`
}

// Reset 清空结构体
func (m *TmallMcDeviceCircleCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTmallMcDeviceCircleCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMcDeviceCircleCheckAPIResponse)
	},
}

// GetTmallMcDeviceCircleCheckAPIResponse 从 sync.Pool 获取 TmallMcDeviceCircleCheckAPIResponse
func GetTmallMcDeviceCircleCheckAPIResponse() *TmallMcDeviceCircleCheckAPIResponse {
	return poolTmallMcDeviceCircleCheckAPIResponse.Get().(*TmallMcDeviceCircleCheckAPIResponse)
}

// ReleaseTmallMcDeviceCircleCheckAPIResponse 将 TmallMcDeviceCircleCheckAPIResponse 保存到 sync.Pool
func ReleaseTmallMcDeviceCircleCheckAPIResponse(v *TmallMcDeviceCircleCheckAPIResponse) {
	v.Reset()
	poolTmallMcDeviceCircleCheckAPIResponse.Put(v)
}
