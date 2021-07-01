package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人信息创建 API返回值 
tmall.servicecenter.worker.create

服务商工人信息创建
*/
type TmallServicecenterWorkerCreateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkerCreateAPIResponseModel
}

// 服务商工人信息创建 成功返回结果
type TmallServicecenterWorkerCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_worker_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
