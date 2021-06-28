package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部离线统计数据上报 APIResponse
aliyun.alink.data.stat.report

外部合作厂商上报设备的明细数据，或者离线统计数据。
*/
type AliyunAlinkDataStatReportAPIResponse struct {
    model.CommonResponse
    // Response *AliyunAlinkDataStatReportResponse `json:"aliyun_alink_data_stat_report_response,omitempty"` 
    AliyunAlinkDataStatReportResponse
}

/* model for simplify = false
type AliyunAlinkDataStatReportResponse struct {

    // 错误消息
    
    Message   string `json:"message,omitempty"`
    

    // 数据入库状态
    
    Module   bool `json:"module,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AliyunAlinkDataStatReportResponse struct {

    // 错误消息
    Message   string `json:"message,omitempty"`

    // 数据入库状态
    Module   bool `json:"module,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
