package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取停开服apk信息 APIResponse
yunos.tvpubadmin.device.apkinfo

获取停开服apk信息
*/
type YunosTvpubadminDeviceApkinfoAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceApkinfoResponse
}

type YunosTvpubadminDeviceApkinfoResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_apkinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *DicControlApkDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
