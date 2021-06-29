package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备Vendor信息 APIResponse
yunos.tvpubadmin.content.device.getvendor

查询设备Vendor信息
*/
type YunosTvpubadminContentDeviceGetvendorAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentDeviceGetvendorResponse
}

type YunosTvpubadminContentDeviceGetvendorResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_device_getvendor_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // list
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
