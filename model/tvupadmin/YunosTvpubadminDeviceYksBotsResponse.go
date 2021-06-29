package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备列表 APIResponse
yunos.tvpubadmin.device.yks.bots

获取设备列表
*/
type YunosTvpubadminDeviceYksBotsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceYksBotsResponse
}

type YunosTvpubadminDeviceYksBotsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_bots_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
