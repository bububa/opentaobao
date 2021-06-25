package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询订单状态 APIRequest
taobao.tvpay.order.query

tv支付查询订单状态
*/
type TaobaoTvpayOrderQueryRequest struct {
    model.Params

    // 设备id
    deviceId   string 

    // 来源
    from   string 

    // 业务订单号
    bizOrderId   string 

    // 是否淘系
    isTao   bool 

    // 支付宝订单号
    orderNo   string 

    // 订单类型
    orderType   string 

    // 外部订单号
    outOrderNo   string 

    // 牌照方
    license   string 

}

func NewTaobaoTvpayOrderQueryRequest() *TaobaoTvpayOrderQueryRequest{
    return &TaobaoTvpayOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayOrderQueryRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.query"
}

func (r TaobaoTvpayOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayOrderQueryRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayOrderQueryRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayOrderQueryRequest) SetBizOrderId(bizOrderId string) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetBizOrderId() string {
    return r.bizOrderId
}

func (r *TaobaoTvpayOrderQueryRequest) SetIsTao(isTao bool) error {
    r.isTao = isTao
    r.Set("is_tao", isTao)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetIsTao() bool {
    return r.isTao
}

func (r *TaobaoTvpayOrderQueryRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *TaobaoTvpayOrderQueryRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetOrderType() string {
    return r.orderType
}

func (r *TaobaoTvpayOrderQueryRequest) SetOutOrderNo(outOrderNo string) error {
    r.outOrderNo = outOrderNo
    r.Set("out_order_no", outOrderNo)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetOutOrderNo() string {
    return r.outOrderNo
}

func (r *TaobaoTvpayOrderQueryRequest) SetLicense(license string) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r TaobaoTvpayOrderQueryRequest) GetLicense() string {
    return r.license
}

