package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据时间段查询服务商的服务预警消息列表(15分钟内) APIResponse
tmall.servicecenter.servicemonitormessage.search

根据时间段查询服务商的服务预警消息列表(15分钟内)
*/
type TmallServicecenterServicemonitormessageSearchAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServicemonitormessageSearchResponse `json:"tmall_servicecenter_servicemonitormessage_search_response,omitempty"`
}

type TmallServicecenterServicemonitormessageSearchResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
