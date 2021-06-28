package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同城令牌打印接口 APIResponse
alibaba.ifp.fulfill.warehouse.token.query

用于仓内作业打印包裹信息
*/
type AlibabaIfpFulfillWarehouseTokenQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ifp_fulfill_warehouse_token_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    WorkResult   *WorkResult `json:"work_result,omitempty" xml:"