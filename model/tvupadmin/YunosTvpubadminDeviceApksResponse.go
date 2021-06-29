package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服apk列表 APIResponse
yunos.tvpubadmin.device.apks

获取停开服apk列表
*/
type YunosTvpubadminDeviceApksAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceApksResponse
}

type YunosTvpubadminDeviceApksResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_apks_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    ApkList   []DicControlApkDO `json:"apk_list,omitempty" xml:"apk_list>dic_control_apk_do,omitempty"`
    
    
}
