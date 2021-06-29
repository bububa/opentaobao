package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单延住接口 API请求
taobao.xhotel.order.alipayface.extend

信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新
*/
type TaobaoXhotelOrderAlipayfaceExtendRequest struct {
    model.Params
    // 延住后的离店日期(最多总共入住天数不能超过9间夜)
    _checkOut   string
    // 阿里旅行订单id,和outOrderId必须至少传入一个
    _tid   int64
    // 卖家系统订单号,和tid必须至少传入一个
    _outOrderId   string
    // 延住房费,注意不包含原有的房费金额,单位为分
    _extendFee   int64
    // 延住日期段的每日房价信息,注意不包括原有的日期房价
    _extendDailyPriceInfo   string
}

// 初始化TaobaoXhotelOrderAlipayfaceExtendRequest对象
func NewTaobaoXhotelOrderAlipayfaceExtendRequest() *TaobaoXhotelOrderAlipayfaceExtendRequest{
    return &TaobaoXhotelOrderAlipayfaceExtendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.extend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CheckOut Setter
// 延住后的离店日期(最多总共入住天数不能超过9间夜)
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetCheckOut(_checkOut string) error {
    r._checkOut = _checkOut
    r.Set("check_out", _checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetCheckOut() string {
    return r._checkOut
}
// Tid Setter
// 阿里旅行订单id,和outOrderId必须至少传入一个
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetTid() int64 {
    return r._tid
}
// OutOrderId Setter
// 卖家系统订单号,和tid必须至少传入一个
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetOutOrderId() string {
    return r._outOrderId
}
// ExtendFee Setter
// 延住房费,注意不包含原有的房费金额,单位为分
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetExtendFee(_extendFee int64) error {
    r._extendFee = _extendFee
    r.Set("extend_fee", _extendFee)
    return nil
}

// ExtendFee Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetExtendFee() int64 {
    return r._extendFee
}
// ExtendDailyPriceInfo Setter
// 延住日期段的每日房价信息,注意不包括原有的日期房价
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetExtendDailyPriceInfo(_extendDailyPriceInfo string) error {
    r._extendDailyPriceInfo = _extendDailyPriceInfo
    r.Set("extend_daily_price_info", _extendDailyPriceInfo)
    return nil
}

// ExtendDailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetExtendDailyPriceInfo() string {
    return r._extendDailyPriceInfo
}
