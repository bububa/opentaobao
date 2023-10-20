package cmns

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaDeviceGetAPIResponse 设备详情查询 API返回值
// yunos.service.cmns.coa.device.get
//
// 第三方应用开发者调用此接口查询设备详情
type YunosServiceCmnsCoaDeviceGetAPIResponse struct {
	model.CommonResponse
	YunosServiceCmnsCoaDeviceGetAPIResponseModel
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaDeviceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosServiceCmnsCoaDeviceGetAPIResponseModel).Reset()
}

// YunosServiceCmnsCoaDeviceGetAPIResponseModel is 设备详情查询 成功返回结果
type YunosServiceCmnsCoaDeviceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_device_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备详情
	DeviceList []DeviceResult `json:"device_list,omitempty" xml:"device_list>device_result,omitempty"`
	// 接口查询出错提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 200表示查询成功
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaDeviceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeviceList = m.DeviceList[:0]
	m.Message = ""
	m.Status = 0
}

var poolYunosServiceCmnsCoaDeviceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosServiceCmnsCoaDeviceGetAPIResponse)
	},
}

// GetYunosServiceCmnsCoaDeviceGetAPIResponse 从 sync.Pool 获取 YunosServiceCmnsCoaDeviceGetAPIResponse
func GetYunosServiceCmnsCoaDeviceGetAPIResponse() *YunosServiceCmnsCoaDeviceGetAPIResponse {
	return poolYunosServiceCmnsCoaDeviceGetAPIResponse.Get().(*YunosServiceCmnsCoaDeviceGetAPIResponse)
}

// ReleaseYunosServiceCmnsCoaDeviceGetAPIResponse 将 YunosServiceCmnsCoaDeviceGetAPIResponse 保存到 sync.Pool
func ReleaseYunosServiceCmnsCoaDeviceGetAPIResponse(v *YunosServiceCmnsCoaDeviceGetAPIResponse) {
	v.Reset()
	poolYunosServiceCmnsCoaDeviceGetAPIResponse.Put(v)
}
