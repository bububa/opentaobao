package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流单信息查询 APIResponse
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API
*/
type TmallServicecenterWorkcardLogisticsorderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_logisticsorder_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"