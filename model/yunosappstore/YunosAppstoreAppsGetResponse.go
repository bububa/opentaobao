package yunosappstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据包名列表获取应用信息列表 APIResponse
yunos.appstore.apps.get

根据包名列表获取应用信息列表
*/
type YunosAppstoreAppsGetAPIResponse struct {
    model.CommonResponse
    YunosAppstoreAppsGetResponse
}

type YunosAppstoreAppsGetResponse struct {
    XMLName xml.Name `xml:"yunos_appstore_apps_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 应用信息列表
    
    AppInfoList   []AppInfo `json:"app_info_list,omitempty" xml:"app_info_list>app_info,omitempty"`
    
    
}
