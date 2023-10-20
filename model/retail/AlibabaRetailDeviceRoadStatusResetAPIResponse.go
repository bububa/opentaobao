package retail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceRoadStatusResetAPIResponse 贩卖机货道解锁 API返回值
// alibaba.retail.device.road.status.reset
//
// 贩卖机货道解锁
type AlibabaRetailDeviceRoadStatusResetAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDeviceRoadStatusResetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceRoadStatusResetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailDeviceRoadStatusResetAPIResponseModel).Reset()
}

// AlibabaRetailDeviceRoadStatusResetAPIResponseModel is 贩卖机货道解锁 成功返回结果
type AlibabaRetailDeviceRoadStatusResetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_road_status_reset_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceRoadStatusResetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = false
}

var poolAlibabaRetailDeviceRoadStatusResetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDeviceRoadStatusResetAPIResponse)
	},
}

// GetAlibabaRetailDeviceRoadStatusResetAPIResponse 从 sync.Pool 获取 AlibabaRetailDeviceRoadStatusResetAPIResponse
func GetAlibabaRetailDeviceRoadStatusResetAPIResponse() *AlibabaRetailDeviceRoadStatusResetAPIResponse {
	return poolAlibabaRetailDeviceRoadStatusResetAPIResponse.Get().(*AlibabaRetailDeviceRoadStatusResetAPIResponse)
}

// ReleaseAlibabaRetailDeviceRoadStatusResetAPIResponse 将 AlibabaRetailDeviceRoadStatusResetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailDeviceRoadStatusResetAPIResponse(v *AlibabaRetailDeviceRoadStatusResetAPIResponse) {
	v.Reset()
	poolAlibabaRetailDeviceRoadStatusResetAPIResponse.Put(v)
}
