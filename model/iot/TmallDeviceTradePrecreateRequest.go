package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上预创建天猫订单 APIRequest
tmall.device.trade.precreate

智能硬件上预创建天猫订单。
1，use_open_price不再需要传入，使用unit_price传入单价。
2，订单默认5分钟自动关闭，没有付款的订单在手机淘宝不可见。
3，同一个码只运行一个用户扫码，多个用户扫一个码会报错 订单不存在。
*/
type TmallDeviceTradePrecreateRequest struct {
    model.Params

    // 交易类型。1，售卖。2，派样
    type   int64 

    // 商品列表
    itemList   []TradeItemDo 

    // 设备业务编码
    deviceCode   string 

    // 外部订单ID，用户下保证唯一。最大长度22
    outTradeId   string 

    // 回调地址，当订单创建，付款成功后，会收到回调。必须是https地址，HTTP响应必须包含success，否则系统会重发
    callbackUrl   string 

    // 导购员手机号码
    baPhone   string 

    // 导购员所属门店
    baStoreId   int64 

}

func NewTmallDeviceTradePrecreateRequest() *TmallDeviceTradePrecreateRequest{
    return &TmallDeviceTradePrecreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallDeviceTradePrecreateRequest) GetApiMethodName() string {
    return "tmall.device.trade.precreate"
}

func (r TmallDeviceTradePrecreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallDeviceTradePrecreateRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetType() int64 {
    return r.type
}

func (r *TmallDeviceTradePrecreateRequest) SetItemList(itemList []TradeItemDo) error {
    r.itemList = itemList
    r.Set("item_list", itemList)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetItemList() []TradeItemDo {
    return r.itemList
}

func (r *TmallDeviceTradePrecreateRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TmallDeviceTradePrecreateRequest) SetOutTradeId(outTradeId string) error {
    r.outTradeId = outTradeId
    r.Set("out_trade_id", outTradeId)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetOutTradeId() string {
    return r.outTradeId
}

func (r *TmallDeviceTradePrecreateRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

func (r *TmallDeviceTradePrecreateRequest) SetBaPhone(baPhone string) error {
    r.baPhone = baPhone
    r.Set("ba_phone", baPhone)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetBaPhone() string {
    return r.baPhone
}

func (r *TmallDeviceTradePrecreateRequest) SetBaStoreId(baStoreId int64) error {
    r.baStoreId = baStoreId
    r.Set("ba_store_id", baStoreId)
    return nil
}

func (r TmallDeviceTradePrecreateRequest) GetBaStoreId() int64 {
    return r.baStoreId
}

