package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销订单查询订单创建结果 API请求
alitrip.xhotel.channel.order.create.res.query

针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
*/
type AlitripXhotelChannelOrderCreateResQueryRequest struct {
    model.Params
    // 外部渠道订单号
    _outSourceOrderId   string
}

// 初始化AlitripXhotelChannelOrderCreateResQueryRequest对象
func NewAlitripXhotelChannelOrderCreateResQueryRequest() *AlitripXhotelChannelOrderCreateResQueryRequest{
    return &AlitripXhotelChannelOrderCreateResQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderCreateResQueryRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.order.create.res.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderCreateResQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutSourceOrderId Setter
// 外部渠道订单号
func (r *AlitripXhotelChannelOrderCreateResQueryRequest) SetOutSourceOrderId(_outSourceOrderId string) error {
    r._outSourceOrderId = _outSourceOrderId
    r.Set("out_source_order_id", _outSourceOrderId)
    return nil
}

// OutSourceOrderId Getter
func (r AlitripXhotelChannelOrderCreateResQueryRequest) GetOutSourceOrderId() string {
    return r._outSourceOrderId
}
