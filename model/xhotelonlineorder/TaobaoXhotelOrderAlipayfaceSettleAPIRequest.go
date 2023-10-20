package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderalipayfacesettleAPIRequest 信用住订单结账接口 API请求
// taobao.xhotel.order.alipayface.settle
//
// 用于离店付订单在客人离店后，发起结账以及扣款等后续动作
type TaobaoxhotelorderalipayfacesettleAPIRequest struct {
	model.Params
	// 单间房明细
	_roomSettleInfoList []RoomSettleInfo
	// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
	_otherFeeDetail string
	// 商家订单号
	_outId string
	// 入住房间号
	_roomNo string
	// 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
	_dailyPriceInfo string
	// 实际离店日期，用于校验总房费收取
	_checkOut string
	// 备注
	_memo string
	// 币种标识，默认人民币
	_currencyCode string
	// 汇率
	_currencyRate string
	// 酒店外部编码
	_hotelCode string
	// 淘宝订单id,必须填写
	_tid int64
	// 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
	_totalRoomFee int64
	// 杂费总额(不能为负数)
	_otherFee int64
	// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
	_containGuarantee int64
	// 结账变价标识，0未变价，1变价
	_priceChange int64
	// 税和服务费
	_taxAndFee int64
	// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
	_amount int64
}

// NewTaobaoxhotelorderalipayfacesettleRequest 初始化TaobaoxhotelorderalipayfacesettleAPIRequest对象
func NewTaobaoxhotelorderalipayfacesettleRequest() *TaobaoxhotelorderalipayfacesettleAPIRequest {
	return &TaobaoxhotelorderalipayfacesettleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.settle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoomSettleInfoList is RoomSettleInfoList Setter
// 单间房明细
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetRoomSettleInfoList(_roomSettleInfoList []RoomSettleInfo) error {
	r._roomSettleInfoList = _roomSettleInfoList
	r.Set("room_settle_info_list", _roomSettleInfoList)
	return nil
}

// GetRoomSettleInfoList RoomSettleInfoList Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetRoomSettleInfoList() []RoomSettleInfo {
	return r._roomSettleInfoList
}

// SetOtherFeeDetail is OtherFeeDetail Setter
// 杂费明细,如果otherFee&gt;0则该字段必须设置,并和杂费金额相吻合
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetOtherFeeDetail(_otherFeeDetail string) error {
	r._otherFeeDetail = _otherFeeDetail
	r.Set("other_fee_detail", _otherFeeDetail)
	return nil
}

// GetOtherFeeDetail OtherFeeDetail Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetOtherFeeDetail() string {
	return r._otherFeeDetail
}

// SetOutId is OutId Setter
// 商家订单号
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetOutId() string {
	return r._outId
}

// SetRoomNo is RoomNo Setter
// 入住房间号
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetRoomNo(_roomNo string) error {
	r._roomNo = _roomNo
	r.Set("room_no", _roomNo)
	return nil
}

// GetRoomNo RoomNo Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetRoomNo() string {
	return r._roomNo
}

// SetDailyPriceInfo is DailyPriceInfo Setter
// 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
	r._dailyPriceInfo = _dailyPriceInfo
	r.Set("daily_price_info", _dailyPriceInfo)
	return nil
}

// GetDailyPriceInfo DailyPriceInfo Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetDailyPriceInfo() string {
	return r._dailyPriceInfo
}

// SetCheckOut is CheckOut Setter
// 实际离店日期，用于校验总房费收取
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetMemo is Memo Setter
// 备注
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetMemo() string {
	return r._memo
}

// SetCurrencyCode is CurrencyCode Setter
// 币种标识，默认人民币
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetCurrencyRate is CurrencyRate Setter
// 汇率
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetCurrencyRate(_currencyRate string) error {
	r._currencyRate = _currencyRate
	r.Set("currency_rate", _currencyRate)
	return nil
}

// GetCurrencyRate CurrencyRate Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetCurrencyRate() string {
	return r._currencyRate
}

// SetHotelCode is HotelCode Setter
// 酒店外部编码
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetTid is Tid Setter
// 淘宝订单id,必须填写
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetTid() int64 {
	return r._tid
}

// SetTotalRoomFee is TotalRoomFee Setter
// 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetTotalRoomFee(_totalRoomFee int64) error {
	r._totalRoomFee = _totalRoomFee
	r.Set("total_room_fee", _totalRoomFee)
	return nil
}

// GetTotalRoomFee TotalRoomFee Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetTotalRoomFee() int64 {
	return r._totalRoomFee
}

// SetOtherFee is OtherFee Setter
// 杂费总额(不能为负数)
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetOtherFee(_otherFee int64) error {
	r._otherFee = _otherFee
	r.Set("other_fee", _otherFee)
	return nil
}

// GetOtherFee OtherFee Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetOtherFee() int64 {
	return r._otherFee
}

// SetContainGuarantee is ContainGuarantee Setter
// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetContainGuarantee(_containGuarantee int64) error {
	r._containGuarantee = _containGuarantee
	r.Set("contain_guarantee", _containGuarantee)
	return nil
}

// GetContainGuarantee ContainGuarantee Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetContainGuarantee() int64 {
	return r._containGuarantee
}

// SetPriceChange is PriceChange Setter
// 结账变价标识，0未变价，1变价
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetPriceChange(_priceChange int64) error {
	r._priceChange = _priceChange
	r.Set("price_change", _priceChange)
	return nil
}

// GetPriceChange PriceChange Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetPriceChange() int64 {
	return r._priceChange
}

// SetTaxAndFee is TaxAndFee Setter
// 税和服务费
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetTaxAndFee(_taxAndFee int64) error {
	r._taxAndFee = _taxAndFee
	r.Set("tax_and_fee", _taxAndFee)
	return nil
}

// GetTaxAndFee TaxAndFee Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetTaxAndFee() int64 {
	return r._taxAndFee
}

// SetAmount is Amount Setter
// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
func (r *TaobaoxhotelorderalipayfacesettleAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r TaobaoxhotelorderalipayfacesettleAPIRequest) GetAmount() int64 {
	return r._amount
}
