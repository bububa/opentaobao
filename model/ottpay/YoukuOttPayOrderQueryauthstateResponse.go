package ottpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询连包签约状态 APIResponse
youku.ott.pay.order.queryauthstate

查询CP用户连包商品签约状态
*/
type YoukuOttPayOrderQueryauthstateAPIResponse struct {
    model.CommonResponse
    YoukuOttPayOrderQueryauthstateResponse
}

type YoukuOttPayOrderQueryauthstateResponse struct {
    XMLName xml.Name `xml:"youku_ott_pay_order_queryauthstate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
