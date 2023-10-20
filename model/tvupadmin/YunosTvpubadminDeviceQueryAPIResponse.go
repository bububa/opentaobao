package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceQueryAPIResponse 获取设备列表 API返回值
// yunos.tvpubadmin.device.query
//
// 获取设备列表
type YunosTvpubadminDeviceQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceQueryAPIResponseModel
}

// YunosTvpubadminDeviceQueryAPIResponseModel is 获取设备列表 成功返回结果
type YunosTvpubadminDeviceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}
