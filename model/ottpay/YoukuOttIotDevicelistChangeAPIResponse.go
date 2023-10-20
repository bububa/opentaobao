package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttIotDevicelistChangeAPIResponse iot设备列表变化接口 API返回值
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
type YoukuOttIotDevicelistChangeAPIResponse struct {
	model.CommonResponse
	YoukuOttIotDevicelistChangeAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttIotDevicelistChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttIotDevicelistChangeAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *YoukuOttIotDevicelistChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolYoukuOttIotDevicelistChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttIotDevicelistChangeAPIResponse)
	},
}

// GetYoukuOttIotDevicelistChangeAPIResponse 从 sync.Pool 获取 YoukuOttIotDevicelistChangeAPIResponse
func GetYoukuOttIotDevicelistChangeAPIResponse() *YoukuOttIotDevicelistChangeAPIResponse {
	return poolYoukuOttIotDevicelistChangeAPIResponse.Get().(*YoukuOttIotDevicelistChangeAPIResponse)
}

// ReleaseYoukuOttIotDevicelistChangeAPIResponse 将 YoukuOttIotDevicelistChangeAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttIotDevicelistChangeAPIResponse(v *YoukuOttIotDevicelistChangeAPIResponse) {
	v.Reset()
	poolYoukuOttIotDevicelistChangeAPIResponse.Put(v)
}
