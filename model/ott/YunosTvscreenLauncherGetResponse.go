package ott

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
一体机桌面 APIResponse
yunos.tvscreen.launcher.get

LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务
*/
type YunosTvscreenLauncherGetAPIResponse struct {
    model.CommonResponse
    YunosTvscreenLauncherGetResponse
}

type YunosTvscreenLauncherGetResponse struct {
    XMLName xml.Name `xml:"yunos_tvscreen_launcher_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // headers
    
    Headers   string `json:"headers,omitempty" xml:"headers,omitempty"`

    
    // model
    
    Model   *LauncherDo `json:"model,omitempty" xml:"model,omitempty"`

    
    // 状态码
    
    HttpStatusCode   int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`

    
    // 业务扩展
    
    BizExtMap   string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`

    
}
