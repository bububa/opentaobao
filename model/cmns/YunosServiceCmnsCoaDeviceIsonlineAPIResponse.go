package cmns

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaDeviceIsonlineAPIResponse 根据设备id查询设备是否在线 API返回值
// yunos.service.cmns.coa.device.isonline
//
// 根据设备id查询设备是否在线
type YunosServiceCmnsCoaDeviceIsonlineAPIResponse struct {
	model.CommonResponse
	YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaDeviceIsonlineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel).Reset()
}

// YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel is 根据设备id查询设备是否在线 成功返回结果
type YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_device_isonline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaDeviceIsonlineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Data = 0
	m.Status = 0
}

var poolYunosServiceCmnsCoaDeviceIsonlineAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosServiceCmnsCoaDeviceIsonlineAPIResponse)
	},
}

// GetYunosServiceCmnsCoaDeviceIsonlineAPIResponse 从 sync.Pool 获取 YunosServiceCmnsCoaDeviceIsonlineAPIResponse
func GetYunosServiceCmnsCoaDeviceIsonlineAPIResponse() *YunosServiceCmnsCoaDeviceIsonlineAPIResponse {
	return poolYunosServiceCmnsCoaDeviceIsonlineAPIResponse.Get().(*YunosServiceCmnsCoaDeviceIsonlineAPIResponse)
}

// ReleaseYunosServiceCmnsCoaDeviceIsonlineAPIResponse 将 YunosServiceCmnsCoaDeviceIsonlineAPIResponse 保存到 sync.Pool
func ReleaseYunosServiceCmnsCoaDeviceIsonlineAPIResponse(v *YunosServiceCmnsCoaDeviceIsonlineAPIResponse) {
	v.Reset()
	poolYunosServiceCmnsCoaDeviceIsonlineAPIResponse.Put(v)
}
