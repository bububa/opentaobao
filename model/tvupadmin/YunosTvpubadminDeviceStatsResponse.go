package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备统计数据 API返回值 
yunos.tvpubadmin.device.stats

获取设备统计数据
*/
type YunosTvpubadminDeviceStatsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceStatsResponse
}

// 获取设备统计数据 成功返回结果
type YunosTvpubadminDeviceStatsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_stats_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // object
    Object   *StatsDeviceInfoDO `json:"object,omitempty" xml:"object,omitempty"`
}
