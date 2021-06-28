package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取工人标签 APIResponse
tmall.servicecenter.worker.taglist.get

服务商获取对应工人的标签
*/
type TmallServicecenterWorkerTaglistGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkerTaglistGetResponse `json:"tmall_servicecenter_worker_taglist_get_response,omitempty"` 
    TmallServicecenterWorkerTaglistGetResponse
}

/* model for simplify = false
type TmallServicecenterWorkerTaglistGetResponse struct {

    // 工人的能力标签
    
    Result  *struct {
        WorkerTag  *WorkerTag `json:"worker_tag,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkerTaglistGetResponse struct {

    // 工人的能力标签
    Result   *WorkerTag `json:"result,omitempty"`

}
