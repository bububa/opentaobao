package tvupadmin

import (
	"encoding/xml"

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

// YunosTvpubadminDeviceUpdateappstatusAPIResponseModel is 更新应用版本审核状态 成功返回结果
type YunosTvpubadminDeviceUpdateappstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_updateappstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
