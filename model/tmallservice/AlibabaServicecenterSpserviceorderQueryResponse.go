package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单查询 APIResponse
alibaba.servicecenter.spserviceorder.query

服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
*/
type AlibabaServicecenterSpserviceorderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_servicecenter_spserviceorder_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求结果
    
    Result   *AlibabaServicecenterSpserviceorderQueryResult `json:"result,omitempty" xml:"