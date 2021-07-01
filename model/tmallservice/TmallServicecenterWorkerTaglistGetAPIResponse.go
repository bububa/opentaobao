package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取工人标签 API返回值 
tmall.servicecenter.worker.taglist.get

服务商获取对应工人的标签
*/
type TmallServicecenterWorkerTaglistGetAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkerTaglistGetAPIResponseModel
}

// 获取工人标签 成功返回结果
type TmallServicecenterWorkerTaglistGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_worker_taglist_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 工人的能力标签
    Result   *WorkerTag `json:"result,omitempty" xml:"result,omitempty"`
}
