package gameact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeactivitymachineidgetAPIResponse 获取设备号 API返回值
// taobao.de.activity.machineid.get
//
// 获取机器设备id
type TaobaodeactivitymachineidgetAPIResponse struct {
	model.CommonResponse
	TaobaodeactivitymachineidgetAPIResponseModel
}

// TaobaodeactivitymachineidgetAPIResponseModel is 获取设备号 成功返回结果
type TaobaodeactivitymachineidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"de_activity_machineid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 机器号
	MachineId string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`
}
