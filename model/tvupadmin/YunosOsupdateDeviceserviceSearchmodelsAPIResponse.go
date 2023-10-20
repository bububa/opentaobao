package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateDeviceserviceSearchmodelsAPIResponse 根据关键词检索设备型号 API返回值
// yunos.osupdate.deviceservice.searchmodels
//
// 根据关键词检索设备型号
type YunosOsupdateDeviceserviceSearchmodelsAPIResponse struct {
	model.CommonResponse
	YunosOsupdateDeviceserviceSearchmodelsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateDeviceserviceSearchmodelsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateDeviceserviceSearchmodelsAPIResponseModel).Reset()
}

// YunosOsupdateDeviceserviceSearchmodelsAPIResponseModel is 根据关键词检索设备型号 成功返回结果
type YunosOsupdateDeviceserviceSearchmodelsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_deviceservice_searchmodels_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	ModelList []DeviceEntryDto `json:"model_list,omitempty" xml:"model_list>device_entry_dto,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateDeviceserviceSearchmodelsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModelList = m.ModelList[:0]
}

var poolYunosOsupdateDeviceserviceSearchmodelsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateDeviceserviceSearchmodelsAPIResponse)
	},
}

// GetYunosOsupdateDeviceserviceSearchmodelsAPIResponse 从 sync.Pool 获取 YunosOsupdateDeviceserviceSearchmodelsAPIResponse
func GetYunosOsupdateDeviceserviceSearchmodelsAPIResponse() *YunosOsupdateDeviceserviceSearchmodelsAPIResponse {
	return poolYunosOsupdateDeviceserviceSearchmodelsAPIResponse.Get().(*YunosOsupdateDeviceserviceSearchmodelsAPIResponse)
}

// ReleaseYunosOsupdateDeviceserviceSearchmodelsAPIResponse 将 YunosOsupdateDeviceserviceSearchmodelsAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateDeviceserviceSearchmodelsAPIResponse(v *YunosOsupdateDeviceserviceSearchmodelsAPIResponse) {
	v.Reset()
	poolYunosOsupdateDeviceserviceSearchmodelsAPIResponse.Put(v)
}
