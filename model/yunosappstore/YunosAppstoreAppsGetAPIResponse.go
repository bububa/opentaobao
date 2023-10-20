package yunosappstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosappstoreappsgetAPIResponse 根据包名列表获取应用信息列表 API返回值
// yunos.appstore.apps.get
//
// 根据包名列表获取应用信息列表
type YunosappstoreappsgetAPIResponse struct {
	model.CommonResponse
	YunosappstoreappsgetAPIResponseModel
}

// YunosappstoreappsgetAPIResponseModel is 根据包名列表获取应用信息列表 成功返回结果
type YunosappstoreappsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_apps_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应用信息列表
	AppInfoList []AppInfo `json:"app_info_list,omitempty" xml:"app_info_list>app_info,omitempty"`
}
