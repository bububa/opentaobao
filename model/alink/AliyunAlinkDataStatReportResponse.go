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
    AliyunAlinkDataStatReportResponse
}

type AliyunAlinkDataStatReportResponse struct {
    XMLName xml.Name `xml:"aliyun_alink_data_stat_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 数据入库状态
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
    // 状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`

    
    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
