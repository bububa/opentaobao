package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 APIResponse
yunos.tvpubadmin.device.tvid

通过UUID查询终端信息
*/
type YunosTvpubadminDeviceTvidAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceTvidResponse
}

type YunosTvpubadminDeviceTvidResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_tvid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *DeviceInfoDo `json:"object,omitempty" xml:"object,omitempty"`

    
}
