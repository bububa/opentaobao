package xhotelofficial

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住结账接口 API请求
taobao.xhotel.order.official.settle.put

用于酒店官网信用住商家结账调用
*/
type TaobaoXhotelOrderOfficialSettlePutAPIRequest struct {
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
}

// 初始化TaobaoXhotelOrderOfficialSettlePutAPIRequest对象
func NewTaobaoXhotelOrderOfficialSettlePutRequest() *TaobaoXhotelOrderOfficialSettlePutAPIRequest{
    return &TaobaoXhotelOrderOfficialSettlePutAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.settle.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单id,必须填写
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetTid() int64 {
    return r._tid
}
// TotalRoomFee Setter
// 房费总额(必须大于0)
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetTotalRoomFee(_totalRoomFee int64) error {
    r._totalRoomFee = _totalRoomFee
    r.Set("total_room_fee", _totalRoomFee)
    return nil
}

// TotalRoomFee Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetTotalRoomFee() int64 {
    return r._totalRoomFee
}
// OtherFee Setter
// 杂费总额(不能为负数)
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetOtherFee(_otherFee int64) error {
    r._otherFee = _otherFee
    r.Set("other_fee", _otherFee)
    return nil
}

// OtherFee Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetOtherFee() int64 {
    return r._otherFee
}
// OtherFeeDetail Setter
// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetOtherFeeDetail(_otherFeeDetail string) error {
    r._otherFeeDetail = _otherFeeDetail
    r.Set("other_fee_detail", _otherFeeDetail)
    return nil
}

// OtherFeeDetail Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetOtherFeeDetail() string {
    return r._otherFeeDetail
}
// OutId Setter
// 商家订单号
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetOutId() string {
    return r._outId
}
// RoomNo Setter
// 入住房间号
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetRoomNo(_roomNo string) error {
    r._roomNo = _roomNo
    r.Set("room_no", _roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetRoomNo() string {
    return r._roomNo
}
// DailyPriceInfo Setter
// 每日房价,json格式,如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
    r._dailyPriceInfo = _dailyPriceInfo
    r.Set("daily_price_info", _dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetDailyPriceInfo() string {
    return r._dailyPriceInfo
}
// CheckOut Setter
// 实际离店日期，用于校验总房费收取
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetCheckOut(_checkOut string) error {
    r._checkOut = _checkOut
    r.Set("check_out", _checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetCheckOut() string {
    return r._checkOut
}
// Memo Setter
// 备注
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetMemo() string {
    return r._memo
}
// RoomSettleInfoList Setter
// 房间明细列表
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetRoomSettleInfoList(_roomSettleInfoList []RoomSettleInfo) error {
    r._roomSettleInfoList = _roomSettleInfoList
    r.Set("room_settle_info_list", _roomSettleInfoList)
    return nil
}

// RoomSettleInfoList Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetRoomSettleInfoList() []RoomSettleInfo {
    return r._roomSettleInfoList
}
// ContainGuarantee Setter
// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含（多间房结账必须传入）
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetContainGuarantee(_containGuarantee int64) error {
    r._containGuarantee = _containGuarantee
    r.Set("contain_guarantee", _containGuarantee)
    return nil
}

// ContainGuarantee Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetContainGuarantee() int64 {
    return r._containGuarantee
}
// OutUuid Setter
// 结账请求流水号
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetOutUuid(_outUuid string) error {
    r._outUuid = _outUuid
    r.Set("out_uuid", _outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetOutUuid() string {
    return r._outUuid
}
// NotifyUrl Setter
// 请求结果通知地址（暂时无效，无需传入）
func (r *TaobaoXhotelOrderOfficialSettlePutAPIRequest) SetNotifyUrl(_notifyUrl string) error {
    r._notifyUrl = _notifyUrl
    r.Set("notify_url", _notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialSettlePutAPIRequest) GetNotifyUrl() string {
    return r._notifyUrl
}
