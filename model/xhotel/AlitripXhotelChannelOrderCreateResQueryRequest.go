package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销订单查询订单创建结果 APIRequest
alitrip.xhotel.channel.order.create.res.query

针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
*/
type AlitripXhotelChannelOrderCreateResQueryRequest struct {
    model.Params

    // 外部渠道订单号
    outSourceOrderId   string 

}

func NewAlitripXhotelChannelOrderCreateResQueryRequest() *AlitripXhotelChannelOrderCreateResQueryRequest{
    return &AlitripXhotelChannelOrderCreateResQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripXhotelChannelOrderCreateResQueryRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.order.create.res.query"
}

func (r AlitripXhotelChannelOrderCreateResQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripXhotelChannelOrderCreateResQueryRequest) SetOutSourceOrderId(outSourceOrderId string) error {
    r.outSourceOrderId = outSourceOrderId
    r.Set("out_source_order_id", outSourceOrderId)
    return nil
}

func (r AlitripXhotelChannelOrderCreateResQueryRequest) GetOutSourceOrderId() string {
    return r.outSourceOrderId
}

