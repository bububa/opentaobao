package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道分销创建订单接口 API请求
alitrip.xhotel.channel.order.create

创建订单接口服务（如菲住等其他渠道分销提供）
*/
type AlitripXhotelChannelOrderCreateRequest struct {
    model.Params
    // 创建订单参数
    outSourceOrderCreateReq   *OutSourceOrderCreateReq
}

// 初始化AlitripXhotelChannelOrderCreateRequest对象
func NewAlitripXhotelChannelOrderCreateRequest() *AlitripXhotelChannelOrderCreateRequest{
    return &AlitripXhotelChannelOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderCreateRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutSourceOrderCreateReq Setter
// 创建订单参数
func (r *AlitripXhotelChannelOrderCreateRequest) SetOutSourceOrderCreateReq(outSourceOrderCreateReq *OutSourceOrderCreateReq) error {
    r.outSourceOrderCreateReq = outSourceOrderCreateReq
    r.Set("out_source_order_create_req", outSourceOrderCreateReq)
    return nil
}

// OutSourceOrderCreateReq Getter
func (r AlitripXhotelChannelOrderCreateRequest) GetOutSourceOrderCreateReq() *OutSourceOrderCreateReq {
    return r.outSourceOrderCreateReq
}
