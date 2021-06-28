package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务单列表查询 APIResponse
tmall.servicecenter.spserviceorder.query

查询服务单列表
*/
type TmallServicecenterSpserviceorderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_spserviceorder_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回参数
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"