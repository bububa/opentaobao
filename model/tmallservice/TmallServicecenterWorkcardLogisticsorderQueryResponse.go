package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流单信息查询 API返回值 
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API
*/
type TmallServicecenterWorkcardLogisticsorderQueryAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardLogisticsorderQueryResponse
}

// 物流单信息查询 成功返回结果
type TmallServicecenterWorkcardLogisticsorderQueryResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsorder_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
