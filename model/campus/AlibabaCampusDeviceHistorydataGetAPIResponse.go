package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceHistorydataGetAPIResponse 设备历史数据批量获取 API返回值
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
type AlibabaCampusDeviceHistorydataGetAPIResponse struct {
	model.CommonResponse
	AlibabaCampusDeviceHistorydataGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceHistorydataGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusDeviceHistorydataGetAPIResponseModel).Reset()
}

// AlibabaCampusDeviceHistorydataGetAPIResponseModel is 设备历史数据批量获取 成功返回结果
type AlibabaCampusDeviceHistorydataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_device_historydata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusDeviceHistorydataGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusDeviceHistorydataGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusDeviceHistorydataGetAPIResponse)
	},
}

// GetAlibabaCampusDeviceHistorydataGetAPIResponse 从 sync.Pool 获取 AlibabaCampusDeviceHistorydataGetAPIResponse
func GetAlibabaCampusDeviceHistorydataGetAPIResponse() *AlibabaCampusDeviceHistorydataGetAPIResponse {
	return poolAlibabaCampusDeviceHistorydataGetAPIResponse.Get().(*AlibabaCampusDeviceHistorydataGetAPIResponse)
}

// ReleaseAlibabaCampusDeviceHistorydataGetAPIResponse 将 AlibabaCampusDeviceHistorydataGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusDeviceHistorydataGetAPIResponse(v *AlibabaCampusDeviceHistorydataGetAPIResponse) {
	v.Reset()
	poolAlibabaCampusDeviceHistorydataGetAPIResponse.Put(v)
}
