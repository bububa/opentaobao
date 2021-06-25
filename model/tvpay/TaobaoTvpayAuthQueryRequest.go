package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付授权查询 APIRequest
taobao.tvpay.auth.query

查询该用户在指定设备上是否有支付授权
*/
type TaobaoTvpayAuthQueryRequest struct {
    model.Params

    // 设备号
    deviceId   string 

    // 来源
    from   string 

    // 业务订单号
    bizOrderId   string 

    // 是否淘系
    isTao   bool 

    // 支付宝订单号
    orderNo   string 

    // 外部订单号
    outOrderNo   string 

}

func NewTaobaoTvpayAuthQueryRequest() *TaobaoTvpayAuthQueryRequest{
    return &TaobaoTvpayAuthQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayAuthQueryRequest) GetApiMethodName() string {
    return "taobao.tvpay.auth.query"
}

func (r TaobaoTvpayAuthQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayAuthQueryRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayAuthQueryRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayAuthQueryRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayAuthQueryRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayAuthQueryRequest) SetBizOrderId(bizOrderId string) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoTvpayAuthQueryRequest) GetBizOrderId() string {
    return r.bizOrderId
}

func (r *TaobaoTvpayAuthQueryRequest) SetIsTao(isTao bool) error {
    r.isTao = isTao
    r.Set("is_tao", isTao)
    return nil
}

func (r TaobaoTvpayAuthQueryRequest) GetIsTao() bool {
    return r.isTao
}

func (r *TaobaoTvpayAuthQueryRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r TaobaoTvpayAuthQueryRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *TaobaoTvpayAuthQueryRequest) SetOutOrderNo(outOrderNo string) error {
    r.outOrderNo = outOrderNo
    r.Set("out_order_no", outOrderNo)
    return nil
}

func (r TaobaoTvpayAuthQueryRequest) GetOutOrderNo() string {
    return r.outOrderNo
}

