package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据时间段查询服务商的服务预警消息列表(15分钟内) APIResponse
tmall.servicecenter.servicemonitormessage.search

根据时间段查询服务商的服务预警消息列表(15分钟内)
*/
type TmallServicecenterServicemonitormessageSearchAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicemonitormessageSearchResponse
}

type TmallServicecenterServicemonitormessageSearchResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicemonitormessage_search_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
