package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新工单状态 APIResponse
alibaba.westcrm.job.status.update

更新工单处理状态
*/
type AlibabaWestcrmJobStatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmJobStatusUpdateResponse
}

type AlibabaWestcrmJobStatusUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_job_status_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
