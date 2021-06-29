package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新系统版本审核状态 APIResponse
yunos.tvpubadmin.device.updateosstatus

更新系统版本审核状态
*/
type YunosTvpubadminDeviceUpdateosstatusAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceUpdateosstatusResponse
}

type YunosTvpubadminDeviceUpdateosstatusResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_updateosstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
