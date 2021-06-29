package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单结账接口 API请求
taobao.xhotel.order.alipayface.settle

用于离店付订单在客人离店后，发起结账以及扣款等后续动作
*/
type TaobaoXhotelOrderAlipayfaceSettleRequest struct {
    model.Params
    // 备注
    _memo   string
    // 商家订单号
    _outId   string
    // 入住房间号
    _roomNo   string
    // 杂费总额(不能为负数)
    _otherFee   int64
    // 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
    _totalRoomFee   int64
    // 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
    _dailyPriceInfo   string
    // 实际离店日期，用于校验总房费收取
    _checkOut   string
    // 淘宝订单id,必须填写
    _tid   int64
    // 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
    _otherFeeDetail   string
    // 单间房明细
    _roomSettleInfoList   []RoomSettleInfo
    // 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
    _containGuarantee   int64
    // 结账变价标识，0未变价，1变价
    _priceChange   int64
    // 币种标识，默认人民币
    _currencyCode   string
    // 汇率
    _currencyRate   string
    // 税和服务费
    _taxAndFee   int64
    // 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
    _amount   int64
    // 酒店外部编码
    _hotelCode   string
}

// 初始化TaobaoXhotelOrderAlipayfaceSettleRequest对象
func NewTaobaoXhotelOrderAlipayfaceSettleRequest() *TaobaoXhotelOrderAlipayfaceSettleRequest{
    return &TaobaoXhotelOrderAlipayfaceSettleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.settle"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Memo Setter
// 备注
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetMemo() string {
    return r._memo
}
// OutId Setter
// 商家订单号
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetOutId() string {
    return r._outId
}
// RoomNo Setter
// 入住房间号
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetRoomNo(_roomNo string) error {
    r._roomNo = _roomNo
    r.Set("room_no", _roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetRoomNo() string {
    return r._roomNo
}
// OtherFee Setter
// 杂费总额(不能为负数)
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetOtherFee(_otherFee int64) error {
    r._otherFee = _otherFee
    r.Set("other_fee", _otherFee)
    return nil
}

// OtherFee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetOtherFee() int64 {
    return r._otherFee
}
// TotalRoomFee Setter
// 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetTotalRoomFee(_totalRoomFee int64) error {
    r._totalRoomFee = _totalRoomFee
    r.Set("total_room_fee", _totalRoomFee)
    return nil
}

// TotalRoomFee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetTotalRoomFee() int64 {
    return r._totalRoomFee
}
// DailyPriceInfo Setter
// 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
    r._dailyPriceInfo = _dailyPriceInfo
    r.Set("daily_price_info", _dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetDailyPriceInfo() string {
    return r._dailyPriceInfo
}
// CheckOut Setter
// 实际离店日期，用于校验总房费收取
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetCheckOut(_checkOut string) error {
    r._checkOut = _checkOut
    r.Set("check_out", _checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetCheckOut() string {
    return r._checkOut
}
// Tid Setter
// 淘宝订单id,必须填写
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetTid() int64 {
    return r._tid
}
// OtherFeeDetail Setter
// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetOtherFeeDetail(_otherFeeDetail string) error {
    r._otherFeeDetail = _otherFeeDetail
    r.Set("other_fee_detail", _otherFeeDetail)
    return nil
}

// OtherFeeDetail Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetOtherFeeDetail() string {
    return r._otherFeeDetail
}
// RoomSettleInfoList Setter
// 单间房明细
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetRoomSettleInfoList(_roomSettleInfoList []RoomSettleInfo) error {
    r._roomSettleInfoList = _roomSettleInfoList
    r.Set("room_settle_info_list", _roomSettleInfoList)
    return nil
}

// RoomSettleInfoList Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetRoomSettleInfoList() []RoomSettleInfo {
    return r._roomSettleInfoList
}
// ContainGuarantee Setter
// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetContainGuarantee(_containGuarantee int64) error {
    r._containGuarantee = _containGuarantee
    r.Set("contain_guarantee", _containGuarantee)
    return nil
}

// ContainGuarantee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetContainGuarantee() int64 {
    return r._containGuarantee
}
// PriceChange Setter
// 结账变价标识，0未变价，1变价
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetPriceChange(_priceChange int64) error {
    r._priceChange = _priceChange
    r.Set("price_change", _priceChange)
    return nil
}

// PriceChange Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetPriceChange() int64 {
    return r._priceChange
}
// CurrencyCode Setter
// 币种标识，默认人民币
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetCurrencyCode(_currencyCode string) error {
    r._currencyCode = _currencyCode
    r.Set("currency_code", _currencyCode)
    return nil
}

// CurrencyCode Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetCurrencyCode() string {
    return r._currencyCode
}
// CurrencyRate Setter
// 汇率
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetCurrencyRate(_currencyRate string) error {
    r._currencyRate = _currencyRate
    r.Set("currency_rate", _currencyRate)
    return nil
}

// CurrencyRate Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetCurrencyRate() string {
    return r._currencyRate
}
// TaxAndFee Setter
// 税和服务费
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetTaxAndFee(_taxAndFee int64) error {
    r._taxAndFee = _taxAndFee
    r.Set("tax_and_fee", _taxAndFee)
    return nil
}

// TaxAndFee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetTaxAndFee() int64 {
    return r._taxAndFee
}
// Amount Setter
// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetAmount(_amount int64) error {
    r._amount = _amount
    r.Set("amount", _amount)
    return nil
}

// Amount Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetAmount() int64 {
    return r._amount
}
// HotelCode Setter
// 酒店外部编码
func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetHotelCode() string {
    return r._hotelCode
}
