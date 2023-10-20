package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse 获取五道口设备管理信息 API返回值
// taobao.wdk.equipment.deviceadmin.deviceinfo.get
//
// 通过仓编码获取五道口设备管理信息
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponseModel).Reset()
}

// TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponseModel is 获取五道口设备管理信息 成功返回结果
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_deviceadmin_deviceinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoWdkEquipmentDeviceadminDeviceinfoGetHmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse)
	},
}

// GetTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse 从 sync.Pool 获取 TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse
func GetTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse() *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse {
	return poolTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse.Get().(*TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse)
}

// ReleaseTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse 将 TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse(v *TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIResponse.Put(v)
}
