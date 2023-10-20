package yunosappstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstoreAppsGetAPIResponse 根据包名列表获取应用信息列表 API返回值
// yunos.appstore.apps.get
//
// 根据包名列表获取应用信息列表
type YunosAppstoreAppsGetAPIResponse struct {
	model.CommonResponse
	YunosAppstoreAppsGetAPIResponseModel
}

// Reset 清空结构体
func (m *YunosAppstoreAppsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosAppstoreAppsGetAPIResponseModel).Reset()
}

// YunosAppstoreAppsGetAPIResponseModel is 根据包名列表获取应用信息列表 成功返回结果
type YunosAppstoreAppsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_apps_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 应用信息列表
	AppInfoList []AppInfo `json:"app_info_list,omitempty" xml:"app_info_list>app_info,omitempty"`
}

// Reset 清空结构体
func (m *YunosAppstoreAppsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AppInfoList = m.AppInfoList[:0]
}

var poolYunosAppstoreAppsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosAppstoreAppsGetAPIResponse)
	},
}

// GetYunosAppstoreAppsGetAPIResponse 从 sync.Pool 获取 YunosAppstoreAppsGetAPIResponse
func GetYunosAppstoreAppsGetAPIResponse() *YunosAppstoreAppsGetAPIResponse {
	return poolYunosAppstoreAppsGetAPIResponse.Get().(*YunosAppstoreAppsGetAPIResponse)
}

// ReleaseYunosAppstoreAppsGetAPIResponse 将 YunosAppstoreAppsGetAPIResponse 保存到 sync.Pool
func ReleaseYunosAppstoreAppsGetAPIResponse(v *YunosAppstoreAppsGetAPIResponse) {
	v.Reset()
	poolYunosAppstoreAppsGetAPIResponse.Put(v)
}
