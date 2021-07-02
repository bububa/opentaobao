package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceAppupgradequeryAPIResponse 应用升级查询 API返回值
// yunos.tvpubadmin.device.appupgradequery
//
// 应用升级查询
type YunosTvpubadminDeviceAppupgradequeryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceAppupgradequeryAPIResponseModel
}

// YunosTvpubadminDeviceAppupgradequeryAPIResponseModel is 应用升级查询 成功返回结果
type YunosTvpubadminDeviceAppupgradequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_appupgradequery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 具体的数据结构
	ObjectList *PaginationDo `json:"object_list,omitempty" xml:"object_list,omitempty"`
}
