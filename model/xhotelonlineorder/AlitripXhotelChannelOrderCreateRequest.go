package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销创建订单接口 APIRequest
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供）
*/
type AlitripXhotelChannelOrderCreateRequest struct {
    model.Params

    // 创建订单参数
    outSourceOrderCreateReq   *OutSourceOrderCreateReq 

}

func NewAlitripXhotelChannelOrderCreateRequest() *AlitripXhotelChannelOrderCreateRequest{
    return &AlitripXhotelChannelOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripXhotelChannelOrderCreateRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.order.create"
}

func (r AlitripXhotelChannelOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripXhotelChannelOrderCreateRequest) SetOutSourceOrderCreateReq(outSourceOrderCreateReq *OutSourceOrderCreateReq) error {
    r.outSourceOrderCreateReq = outSourceOrderCreateReq
    r.Set("out_source_order_create_req", outSourceOrderCreateReq)
    return nil
}

func (r AlitripXhotelChannelOrderCreateRequest) GetOutSourceOrderCreateReq() *OutSourceOrderCreateReq {
    return r.outSourceOrderCreateReq
}

