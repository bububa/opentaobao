package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人信息创建 APIResponse
tmall.servicecenter.worker.create

服务商工人信息创建
*/
type TmallServicecenterWorkerCreateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkerCreateResponse
}

type TmallServicecenterWorkerCreateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_worker_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
