package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资金流水查询 APIResponse
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alicom_order_preauthorize_query_fund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"