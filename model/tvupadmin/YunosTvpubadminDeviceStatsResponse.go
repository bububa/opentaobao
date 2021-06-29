package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备统计数据 APIResponse
yunos.tvpubadmin.device.stats

获取设备统计数据
*/
type YunosTvpubadminDeviceStatsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceStatsResponse
}

type YunosTvpubadminDeviceStatsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_stats_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   *StatsDeviceInfoDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
