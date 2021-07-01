package xhotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotel"
)

/* 
分销订单查询订单创建结果 
alitrip.xhotel.channel.order.create.res.query

针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
*/
func AlitripXhotelChannelOrderCreateResQuery(clt *core.SDKClient, req *xhotel.AlitripXhotelChannelOrderCreateResQueryAPIRequest, session string) (*xhotel.AlitripXhotelChannelOrderCreateResQueryAPIResponse, error) {
    var resp xhotel.AlitripXhotelChannelOrderCreateResQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
