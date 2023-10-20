package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceUpdateappstatusAPIResponse 更新应用版本审核状态 API返回值
// yunos.tvpubadmin.device.updateappstatus
//
// 更新应用版本审核状态
type YunosTvpubadminDeviceUpdateappstatusAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceUpdateappstatusAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceUpdateappstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceUpdateappstatusAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceUpdateappstatusAPIResponseModel is 更新应用版本审核状态 成功返回结果
type YunosTvpubadminDeviceUpdateappstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_updateappstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceUpdateappstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminDeviceUpdateappstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceUpdateappstatusAPIResponse)
	},
}

// GetYunosTvpubadminDeviceUpdateappstatusAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceUpdateappstatusAPIResponse
func GetYunosTvpubadminDeviceUpdateappstatusAPIResponse() *YunosTvpubadminDeviceUpdateappstatusAPIResponse {
	return poolYunosTvpubadminDeviceUpdateappstatusAPIResponse.Get().(*YunosTvpubadminDeviceUpdateappstatusAPIResponse)
}

// ReleaseYunosTvpubadminDeviceUpdateappstatusAPIResponse 将 YunosTvpubadminDeviceUpdateappstatusAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceUpdateappstatusAPIResponse(v *YunosTvpubadminDeviceUpdateappstatusAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceUpdateappstatusAPIResponse.Put(v)
}
