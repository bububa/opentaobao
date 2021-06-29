package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工人列表 APIResponse
tmall.servicecenter.worker.querypage

服务商查询工人列表
*/
type TmallServicecenterWorkerQuerypageAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkerQuerypageResponse
}

type TmallServicecenterWorkerQuerypageResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_worker_querypage_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
