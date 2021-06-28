package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销订单查询订单创建结果 APIResponse
alitrip.xhotel.channel.order.create.res.query

针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
*/
type AlitripXhotelChannelOrderCreateResQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_xhotel_channel_order_create_res_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *HbsResult `json:"result,omitempty" xml:"