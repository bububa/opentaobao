package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新应用版本审核状态 APIResponse
yunos.tvpubadmin.device.updateappstatus

更新应用版本审核状态
*/
type YunosTvpubadminDeviceUpdateappstatusAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceUpdateappstatusResponse
}

type YunosTvpubadminDeviceUpdateappstatusResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_updateappstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
