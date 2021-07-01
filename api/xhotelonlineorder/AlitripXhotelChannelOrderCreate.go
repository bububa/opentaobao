package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
渠道分销创建订单接口 
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供）
*/
func AlitripXhotelChannelOrderCreate(clt *core.SDKClient, req *xhotelonlineorder.AlitripXhotelChannelOrderCreateAPIRequest, session string) (*xhotelonlineorder.AlitripXhotelChannelOrderCreateAPIResponse, error) {
    var resp xhotelonlineorder.AlitripXhotelChannelOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
