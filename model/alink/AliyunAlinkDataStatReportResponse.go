package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部离线统计数据上报 APIResponse
aliyun.alink.data.stat.report

外部合作厂商上报设备的明细数据，或者离线统计数据。
*/
type AliyunAlinkDataStatReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"aliyun_alink_data_stat_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"