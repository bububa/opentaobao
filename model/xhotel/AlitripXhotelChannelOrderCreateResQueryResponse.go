package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销订单查询订单创建结果 APIResponse
alitrip.xhotel.channel.order.create.res.query

针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
*/
type AlitripXhotelChannelOrderCreateResQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlitripXhotelChannelOrderCreateResQueryResponse `json:"alitrip_xhotel_channel_order_create_res_query_response,omitempty"` 
    AlitripXhotelChannelOrderCreateResQueryResponse
}

/* model for simplify = false
type AlitripXhotelChannelOrderCreateResQueryResponse struct {

    // 结果
    
    Result  *struct {
        HbsResult  *HbsResult `json:"hbs_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripXhotelChannelOrderCreateResQueryResponse struct {

    // 结果
    Result   *HbsResult `json:"result,omitempty"`

}
