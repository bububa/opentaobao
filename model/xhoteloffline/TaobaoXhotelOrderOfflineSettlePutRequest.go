package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住结账专用接口 API请求
taobao.xhotel.order.offline.settle.put

线下信用住结账专用接口
*/
type TaobaoXhotelOrderOfflineSettlePutAPIRequest struct {
    model.Params
    // 淘宝订单id,必须填写
    _tid   int64
    // 房费总额(必须大于0)
    _totalRoomFee   int64
    // 杂费总额(不能为负数)
    _otherFee   int64
    // 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
    _otherFeeDetail   string
    // 商家订单号
    _outId   string
    // 入住房间号
    _roomNo   string
    // 每日房价,json格式,如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
    _dailyPriceInfo   string
    // 实际离店日期，用于校验总房费收取
    _checkOut   string
    // 备注
    _memo   string
    // 房间明细列表
    _roomSettleInfoList   []RoomSettleInfo
    // 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含（多间房结账必须传入）
    _containGuarantee   int64
    // 结账请求流水号
    _outUuid   string
    // 请求结果通知地址（暂时无效，无需传入）
    _notifyUrl   string
    // 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额(单位分)
    _amount   int64
    // 商家酒店编码
    _hotelCode   string
    // 系统商标识
    _vendor   string
}

// 初始化TaobaoXhotelOrderOfflineSettlePutAPIRequest对象
func NewTaobaoXhotelOrderOfflineSettlePutRequest() *TaobaoXhotelOrderOfflineSettlePutAPIRequest{
    return &TaobaoXhotelOrderOfflineSettlePutAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.offline.settle.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单id,必须填写
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetTid() int64 {
    return r._tid
}
// TotalRoomFee Setter
// 房费总额(必须大于0)
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetTotalRoomFee(_totalRoomFee int64) error {
    r._totalRoomFee = _totalRoomFee
    r.Set("total_room_fee", _totalRoomFee)
    return nil
}

// TotalRoomFee Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetTotalRoomFee() int64 {
    return r._totalRoomFee
}
// OtherFee Setter
// 杂费总额(不能为负数)
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetOtherFee(_otherFee int64) error {
    r._otherFee = _otherFee
    r.Set("other_fee", _otherFee)
    return nil
}

// OtherFee Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetOtherFee() int64 {
    return r._otherFee
}
// OtherFeeDetail Setter
// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetOtherFeeDetail(_otherFeeDetail string) error {
    r._otherFeeDetail = _otherFeeDetail
    r.Set("other_fee_detail", _otherFeeDetail)
    return nil
}

// OtherFeeDetail Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetOtherFeeDetail() string {
    return r._otherFeeDetail
}
// OutId Setter
// 商家订单号
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetOutId() string {
    return r._outId
}
// RoomNo Setter
// 入住房间号
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetRoomNo(_roomNo string) error {
    r._roomNo = _roomNo
    r.Set("room_no", _roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetRoomNo() string {
    return r._roomNo
}
// DailyPriceInfo Setter
// 每日房价,json格式,如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
    r._dailyPriceInfo = _dailyPriceInfo
    r.Set("daily_price_info", _dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetDailyPriceInfo() string {
    return r._dailyPriceInfo
}
// CheckOut Setter
// 实际离店日期，用于校验总房费收取
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetCheckOut(_checkOut string) error {
    r._checkOut = _checkOut
    r.Set("check_out", _checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetCheckOut() string {
    return r._checkOut
}
// Memo Setter
// 备注
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetMemo() string {
    return r._memo
}
// RoomSettleInfoList Setter
// 房间明细列表
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetRoomSettleInfoList(_roomSettleInfoList []RoomSettleInfo) error {
    r._roomSettleInfoList = _roomSettleInfoList
    r.Set("room_settle_info_list", _roomSettleInfoList)
    return nil
}

// RoomSettleInfoList Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetRoomSettleInfoList() []RoomSettleInfo {
    return r._roomSettleInfoList
}
// ContainGuarantee Setter
// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含（多间房结账必须传入）
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetContainGuarantee(_containGuarantee int64) error {
    r._containGuarantee = _containGuarantee
    r.Set("contain_guarantee", _containGuarantee)
    return nil
}

// ContainGuarantee Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetContainGuarantee() int64 {
    return r._containGuarantee
}
// OutUuid Setter
// 结账请求流水号
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetOutUuid(_outUuid string) error {
    r._outUuid = _outUuid
    r.Set("out_uuid", _outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetOutUuid() string {
    return r._outUuid
}
// NotifyUrl Setter
// 请求结果通知地址（暂时无效，无需传入）
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetNotifyUrl(_notifyUrl string) error {
    r._notifyUrl = _notifyUrl
    r.Set("notify_url", _notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetNotifyUrl() string {
    return r._notifyUrl
}
// Amount Setter
// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额(单位分)
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetAmount(_amount int64) error {
    r._amount = _amount
    r.Set("amount", _amount)
    return nil
}

// Amount Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetAmount() int64 {
    return r._amount
}
// HotelCode Setter
// 商家酒店编码
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetHotelCode() string {
    return r._hotelCode
}
// Vendor Setter
// 系统商标识
func (r *TaobaoXhotelOrderOfflineSettlePutAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderOfflineSettlePutAPIRequest) GetVendor() string {
    return r._vendor
}
