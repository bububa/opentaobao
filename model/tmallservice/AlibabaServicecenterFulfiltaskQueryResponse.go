package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
核销单查询 APIResponse
alibaba.servicecenter.fulfiltask.query

当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
*/
type AlibabaServicecenterFulfiltaskQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_servicecenter_fulfiltask_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaServicecenterFulfiltaskQueryResult `json:"result,omitempty" xml:"