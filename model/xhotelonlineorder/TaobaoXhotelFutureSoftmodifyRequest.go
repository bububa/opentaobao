package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店信息下发 APIRequest
taobao.xhotel.future.softmodify

未来酒店信息下发，包含PMS订单查询和自助入住
*/
type TaobaoXhotelFutureSoftmodifyRequest struct {
    model.Params

    // 超时时长，默认3s
    expireTime   int64 

    // 淘宝订单号
    tid   int64 

    // 外部订单号
    outOrderId   string 

    // 酒店code
    hotelCode   string 

    // 酒店Id
    hid   int64 

    // 请求报文
    context   string 

    // 操作类型
    operateType   string 

    // 请求唯一标识值
    requestId   string 

}

func NewTaobaoXhotelFutureSoftmodifyRequest() *TaobaoXhotelFutureSoftmodifyRequest{
    return &TaobaoXhotelFutureSoftmodifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetApiMethodName() string {
    return "taobao.xhotel.future.softmodify"
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelFutureSoftmodifyRequest) SetExpireTime(expireTime int64) error {
    r.expireTime = expireTime
    r.Set("expire_time", expireTime)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetExpireTime() int64 {
    return r.expireTime
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetOutOrderId() string {
    return r.outOrderId
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetHid() int64 {
    return r.hid
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetContext(context string) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetContext() string {
    return r.context
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetOperateType() string {
    return r.operateType
}

func (r *TaobaoXhotelFutureSoftmodifyRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r TaobaoXhotelFutureSoftmodifyRequest) GetRequestId() string {
    return r.requestId
}

