package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工人信息查询 API返回值 
tmall.servicecenter.worker.query

查询服务商对应的工人信息
*/
type TmallServicecenterWorkerQueryAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkerQueryAPIResponseModel
}

// 工人信息查询 成功返回结果
type TmallServicecenterWorkerQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_worker_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
