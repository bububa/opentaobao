package idleparttime

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职同步日志 APIResponse
alibaba.idle.parttime.synclog

提供给供应商查询的接口
*/
type AlibabaIdleParttimeSynclogAPIResponse struct {
    model.CommonResponse
    AlibabaIdleParttimeSynclogResponse
}

type AlibabaIdleParttimeSynclogResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_parttime_synclog_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaIdleParttimeSynclogResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
