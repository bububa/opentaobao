package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单结账接口 APIRequest
taobao.xhotel.order.alipayface.settle

用于离店付订单在客人离店后，发起结账以及扣款等后续动作
*/
type TaobaoXhotelOrderAlipayfaceSettleRequest struct {
    model.Params

    // 备注
    memo   string 

    // 商家订单号
    outId   string 

    // 入住房间号
    roomNo   string 

    // 杂费总额(不能为负数)
    otherFee   int64 

    // 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
    totalRoomFee   int64 

    // 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
    dailyPriceInfo   string 

    // 实际离店日期，用于校验总房费收取
    checkOut   string 

    // 淘宝订单id,必须填写
    tid   int64 

    // 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
    otherFeeDetail   string 

    // 单间房明细
    roomSettleInfoList   []RoomSettleInfo 

    // 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
    containGuarantee   int64 

    // 结账变价标识，0未变价，1变价
    priceChange   int64 

    // 币种标识，默认人民币
    currencyCode   string 

    // 汇率
    currencyRate   string 

    // 税和服务费
    taxAndFee   int64 

    // 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
    amount   int64 

    // 酒店外部编码
    hotelCode   string 

}

func NewTaobaoXhotelOrderAlipayfaceSettleRequest() *TaobaoXhotelOrderAlipayfaceSettleRequest{
    return &TaobaoXhotelOrderAlipayfaceSettleRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.settle"
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetMemo() string {
    return r.memo
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetRoomNo(roomNo string) error {
    r.roomNo = roomNo
    r.Set("room_no", roomNo)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetRoomNo() string {
    return r.roomNo
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetOtherFee(otherFee int64) error {
    r.otherFee = otherFee
    r.Set("other_fee", otherFee)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetOtherFee() int64 {
    return r.otherFee
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetTotalRoomFee(totalRoomFee int64) error {
    r.totalRoomFee = totalRoomFee
    r.Set("total_room_fee", totalRoomFee)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetTotalRoomFee() int64 {
    return r.totalRoomFee
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetDailyPriceInfo(dailyPriceInfo string) error {
    r.dailyPriceInfo = dailyPriceInfo
    r.Set("daily_price_info", dailyPriceInfo)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetDailyPriceInfo() string {
    return r.dailyPriceInfo
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetCheckOut() string {
    return r.checkOut
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetOtherFeeDetail(otherFeeDetail string) error {
    r.otherFeeDetail = otherFeeDetail
    r.Set("other_fee_detail", otherFeeDetail)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetOtherFeeDetail() string {
    return r.otherFeeDetail
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetRoomSettleInfoList(roomSettleInfoList []RoomSettleInfo) error {
    r.roomSettleInfoList = roomSettleInfoList
    r.Set("room_settle_info_list", roomSettleInfoList)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetRoomSettleInfoList() []RoomSettleInfo {
    return r.roomSettleInfoList
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetContainGuarantee(containGuarantee int64) error {
    r.containGuarantee = containGuarantee
    r.Set("contain_guarantee", containGuarantee)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetContainGuarantee() int64 {
    return r.containGuarantee
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetPriceChange(priceChange int64) error {
    r.priceChange = priceChange
    r.Set("price_change", priceChange)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetPriceChange() int64 {
    return r.priceChange
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetCurrencyCode(currencyCode string) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetCurrencyCode() string {
    return r.currencyCode
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetCurrencyRate(currencyRate string) error {
    r.currencyRate = currencyRate
    r.Set("currency_rate", currencyRate)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetCurrencyRate() string {
    return r.currencyRate
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetTaxAndFee(taxAndFee int64) error {
    r.taxAndFee = taxAndFee
    r.Set("tax_and_fee", taxAndFee)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetTaxAndFee() int64 {
    return r.taxAndFee
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetAmount(amount int64) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetAmount() int64 {
    return r.amount
}

func (r *TaobaoXhotelOrderAlipayfaceSettleRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceSettleRequest) GetHotelCode() string {
    return r.hotelCode
}

