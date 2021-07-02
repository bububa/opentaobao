package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceSettleAPIRequest 信用住订单结账接口 API请求
// taobao.xhotel.order.alipayface.settle
//
// 用于离店付订单在客人离店后，发起结账以及扣款等后续动作
type TaobaoXhotelOrderAlipayfaceSettleAPIRequest struct {
	model.Params
	// 备注
	_memo string
	// 商家订单号
	_outId string
	// 入住房间号
	_roomNo string
	// 杂费总额(不能为负数)
	_otherFee int64
	// 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
	_totalRoomFee int64
	// 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
	_dailyPriceInfo string
	// 实际离店日期，用于校验总房费收取
	_checkOut string
	// 淘宝订单id,必须填写
	_tid int64
	// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
	_otherFeeDetail string
	// 单间房明细
	_roomSettleInfoList []RoomSettleInfo
	// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
	_containGuarantee int64
	// 结账变价标识，0未变价，1变价
	_priceChange int64
	// 币种标识，默认人民币
	_currencyCode string
	// 汇率
	_currencyRate string
	// 税和服务费
	_taxAndFee int64
	// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
	_amount int64
	// 酒店外部编码
	_hotelCode string
}

// NewTaobaoXhotelOrderAlipayfaceSettleRequest 初始化TaobaoXhotelOrderAlipayfaceSettleAPIRequest对象
func NewTaobaoXhotelOrderAlipayfaceSettleRequest() *TaobaoXhotelOrderAlipayfaceSettleAPIRequest {
	return &TaobaoXhotelOrderAlipayfaceSettleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.settle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMemo is Memo Setter
// 备注
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetMemo() string {
	return r._memo
}

// SetOutId is OutId Setter
// 商家订单号
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetOutId() string {
	return r._outId
}

// SetRoomNo is RoomNo Setter
// 入住房间号
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetRoomNo(_roomNo string) error {
	r._roomNo = _roomNo
	r.Set("room_no", _roomNo)
	return nil
}

// GetRoomNo RoomNo Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetRoomNo() string {
	return r._roomNo
}

// SetOtherFee is OtherFee Setter
// 杂费总额(不能为负数)
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetOtherFee(_otherFee int64) error {
	r._otherFee = _otherFee
	r.Set("other_fee", _otherFee)
	return nil
}

// GetOtherFee OtherFee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetOtherFee() int64 {
	return r._otherFee
}

// SetTotalRoomFee is TotalRoomFee Setter
// 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetTotalRoomFee(_totalRoomFee int64) error {
	r._totalRoomFee = _totalRoomFee
	r.Set("total_room_fee", _totalRoomFee)
	return nil
}

// GetTotalRoomFee TotalRoomFee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetTotalRoomFee() int64 {
	return r._totalRoomFee
}

// SetDailyPriceInfo is DailyPriceInfo Setter
// 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
	r._dailyPriceInfo = _dailyPriceInfo
	r.Set("daily_price_info", _dailyPriceInfo)
	return nil
}

// GetDailyPriceInfo DailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetDailyPriceInfo() string {
	return r._dailyPriceInfo
}

// SetCheckOut is CheckOut Setter
// 实际离店日期，用于校验总房费收取
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetTid is Tid Setter
// 淘宝订单id,必须填写
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetTid() int64 {
	return r._tid
}

// SetOtherFeeDetail is OtherFeeDetail Setter
// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetOtherFeeDetail(_otherFeeDetail string) error {
	r._otherFeeDetail = _otherFeeDetail
	r.Set("other_fee_detail", _otherFeeDetail)
	return nil
}

// GetOtherFeeDetail OtherFeeDetail Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetOtherFeeDetail() string {
	return r._otherFeeDetail
}

// SetRoomSettleInfoList is RoomSettleInfoList Setter
// 单间房明细
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetRoomSettleInfoList(_roomSettleInfoList []RoomSettleInfo) error {
	r._roomSettleInfoList = _roomSettleInfoList
	r.Set("room_settle_info_list", _roomSettleInfoList)
	return nil
}

// GetRoomSettleInfoList RoomSettleInfoList Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetRoomSettleInfoList() []RoomSettleInfo {
	return r._roomSettleInfoList
}

// SetContainGuarantee is ContainGuarantee Setter
// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetContainGuarantee(_containGuarantee int64) error {
	r._containGuarantee = _containGuarantee
	r.Set("contain_guarantee", _containGuarantee)
	return nil
}

// GetContainGuarantee ContainGuarantee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetContainGuarantee() int64 {
	return r._containGuarantee
}

// SetPriceChange is PriceChange Setter
// 结账变价标识，0未变价，1变价
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetPriceChange(_priceChange int64) error {
	r._priceChange = _priceChange
	r.Set("price_change", _priceChange)
	return nil
}

// GetPriceChange PriceChange Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetPriceChange() int64 {
	return r._priceChange
}

// SetCurrencyCode is CurrencyCode Setter
// 币种标识，默认人民币
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetCurrencyRate is CurrencyRate Setter
// 汇率
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetCurrencyRate(_currencyRate string) error {
	r._currencyRate = _currencyRate
	r.Set("currency_rate", _currencyRate)
	return nil
}

// GetCurrencyRate CurrencyRate Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetCurrencyRate() string {
	return r._currencyRate
}

// SetTaxAndFee is TaxAndFee Setter
// 税和服务费
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetTaxAndFee(_taxAndFee int64) error {
	r._taxAndFee = _taxAndFee
	r.Set("tax_and_fee", _taxAndFee)
	return nil
}

// GetTaxAndFee TaxAndFee Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetTaxAndFee() int64 {
	return r._taxAndFee
}

// SetAmount is Amount Setter
// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetAmount() int64 {
	return r._amount
}

// SetHotelCode is HotelCode Setter
// 酒店外部编码
func (r *TaobaoXhotelOrderAlipayfaceSettleAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceSettleAPIRequest) GetHotelCode() string {
	return r._hotelCode
}
