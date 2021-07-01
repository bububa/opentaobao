package ottpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttIotDevicelistChangeAPIResponse
iot设备列表变化接口 API返回值
youku.ott.iot.devicelist.change

iot设备列表变化接口 */
type YoukuOttIotDevicelistChangeAPIResponse struct {
	model.CommonResponse
	YoukuOttIotDevicelistChangeAPIResponseModel
}

// YoukuOttIotDevicelistChangeAPIResponseModel is iot设备列表变化接口 成功返回结果
type YoukuOttIotDevicelistChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_iot_devicelist_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
