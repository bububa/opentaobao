package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上预创建天猫订单 API请求
tmall.device.trade.precreate

智能硬件上预创建天猫订单。
1，use_open_price不再需要传入，使用unit_price传入单价。
2，订单默认5分钟自动关闭，没有付款的订单在手机淘宝不可见。
3，同一个码只运行一个用户扫码，多个用户扫一个码会报错 订单不存在。
*/
type TmallDeviceTradePrecreateAPIRequest struct {
    model.Params
    // 交易类型。1，售卖。2，派样
    _type   int64
    // 商品列表
    _itemList   []TradeItemDO
    // 设备业务编码
    _deviceCode   string
    // 外部订单ID，用户下保证唯一。最大长度22
    _outTradeId   string
    // 回调地址，当订单创建，付款成功后，会收到回调。必须是https地址，HTTP响应必须包含success，否则系统会重发
    _callbackUrl   string
    // 导购员手机号码
    _baPhone   string
    // 导购员所属门店
    _baStoreId   int64
}

// 初始化TmallDeviceTradePrecreateAPIRequest对象
func NewTmallDeviceTradePrecreateRequest() *TmallDeviceTradePrecreateAPIRequest{
    return &TmallDeviceTradePrecreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceTradePrecreateAPIRequest) GetApiMethodName() string {
    return "tmall.device.trade.precreate"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceTradePrecreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 交易类型。1，售卖。2，派样
func (r *TmallDeviceTradePrecreateAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetType() int64 {
    return r._type
}
// ItemList Setter
// 商品列表
func (r *TmallDeviceTradePrecreateAPIRequest) SetItemList(_itemList []TradeItemDO) error {
    r._itemList = _itemList
    r.Set("item_list", _itemList)
    return nil
}

// ItemList Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetItemList() []TradeItemDO {
    return r._itemList
}
// DeviceCode Setter
// 设备业务编码
func (r *TmallDeviceTradePrecreateAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// OutTradeId Setter
// 外部订单ID，用户下保证唯一。最大长度22
func (r *TmallDeviceTradePrecreateAPIRequest) SetOutTradeId(_outTradeId string) error {
    r._outTradeId = _outTradeId
    r.Set("out_trade_id", _outTradeId)
    return nil
}

// OutTradeId Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetOutTradeId() string {
    return r._outTradeId
}
// CallbackUrl Setter
// 回调地址，当订单创建，付款成功后，会收到回调。必须是https地址，HTTP响应必须包含success，否则系统会重发
func (r *TmallDeviceTradePrecreateAPIRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// BaPhone Setter
// 导购员手机号码
func (r *TmallDeviceTradePrecreateAPIRequest) SetBaPhone(_baPhone string) error {
    r._baPhone = _baPhone
    r.Set("ba_phone", _baPhone)
    return nil
}

// BaPhone Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetBaPhone() string {
    return r._baPhone
}
// BaStoreId Setter
// 导购员所属门店
func (r *TmallDeviceTradePrecreateAPIRequest) SetBaStoreId(_baStoreId int64) error {
    r._baStoreId = _baStoreId
    r.Set("ba_store_id", _baStoreId)
    return nil
}

// BaStoreId Getter
func (r TmallDeviceTradePrecreateAPIRequest) GetBaStoreId() int64 {
    return r._baStoreId
}
