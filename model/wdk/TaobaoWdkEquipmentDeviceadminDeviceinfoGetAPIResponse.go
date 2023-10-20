package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponse 获取五道口设备管理信息 API返回值
// taobao.wdk.equipment.deviceadmin.deviceinfo.get
//
// 通过仓编码获取五道口设备管理信息
type TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponse struct {
	model.CommonResponse
	TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponseModel
}

// TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponseModel is 获取五道口设备管理信息 成功返回结果
type TaobaowdkequipmentdeviceadmindeviceinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_equipment_deviceadmin_deviceinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaowdkequipmentdeviceadmindeviceinfogetHmResult `json:"result,omitempty" xml:"result,omitempty"`
}
