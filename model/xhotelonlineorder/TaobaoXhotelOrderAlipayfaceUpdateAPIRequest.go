package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店信用住订单状态更新 API请求
taobao.xhotel.order.alipayface.update

完成对信用住或者面付订单的状态的更新。包含订单状态的确认，入离店状态的更新等等。（不适用于预付订单）
*/
type TaobaoXhotelOrderAlipayfaceUpdateAPIRequest struct {
    model.Params
    // 淘宝订单号,必填
    _tid   int64
    // 操作的类型：12.补录确认号,11.多间房确认无房，10.多间房确认有房，8.取消订单(cancel)酒店端发起取消,必须在和买家协商通过的情况下操作,否则有法务风险; 5.买家未入住(noshow),如果该单有担保，会收取买家的担保金额; 3.核实入住(checkIn); 4.核实离店(checkOut); 1.确认无房(直连卖家禁止该操作),2.确认有房(直连卖家禁止该操作)
    _optType   int64
    // 无房原因分类:1.无房, 2.价格变动, 3.买家原因, 4.其它原因,opt_type=1时必填
    _reasonType   int64
    // 无房原因描述:opt_type=1时必填
    _reasonText   string
    // 入住房间号
    _outRoomNumber   string
    // 客人实际入住日期,opt_type=3/4时必填
    _checkinDate   string
    // 客人实际离店日期,opt_type=4时必填
    _checkoutDate   string
    // 客人实际预定房间数
    _rooms   int64
    // 外部订单号
    _outId   string
    // opt_type为10,11启用，多间房订单号列表，逗号间隔
    _tids   string
    // opt_type为11启用，多间房订单取消原因类型，逗号间隔
    _cancelType   int64
    // opt_type为10,11,12启用，真实操作人
    _operator   string
    // opt_type为12, 订单确认号
    _confirmCode   string
    // 是否自助入住
    _selfCheckin   bool
    // 是否把代理直签的订单同步到酒店，Y为同步，N不同步
    _syncToHotel   string
}

// 初始化TaobaoXhotelOrderAlipayfaceUpdateAPIRequest对象
func NewTaobaoXhotelOrderAlipayfaceUpdateRequest() *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest{
    return &TaobaoXhotelOrderAlipayfaceUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单号,必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetTid() int64 {
    return r._tid
}
// OptType Setter
// 操作的类型：12.补录确认号,11.多间房确认无房，10.多间房确认有房，8.取消订单(cancel)酒店端发起取消,必须在和买家协商通过的情况下操作,否则有法务风险; 5.买家未入住(noshow),如果该单有担保，会收取买家的担保金额; 3.核实入住(checkIn); 4.核实离店(checkOut); 1.确认无房(直连卖家禁止该操作),2.确认有房(直连卖家禁止该操作)
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetOptType(_optType int64) error {
    r._optType = _optType
    r.Set("opt_type", _optType)
    return nil
}

// OptType Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetOptType() int64 {
    return r._optType
}
// ReasonType Setter
// 无房原因分类:1.无房, 2.价格变动, 3.买家原因, 4.其它原因,opt_type=1时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetReasonType(_reasonType int64) error {
    r._reasonType = _reasonType
    r.Set("reason_type", _reasonType)
    return nil
}

// ReasonType Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetReasonType() int64 {
    return r._reasonType
}
// ReasonText Setter
// 无房原因描述:opt_type=1时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetReasonText(_reasonText string) error {
    r._reasonText = _reasonText
    r.Set("reason_text", _reasonText)
    return nil
}

// ReasonText Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetReasonText() string {
    return r._reasonText
}
// OutRoomNumber Setter
// 入住房间号
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetOutRoomNumber(_outRoomNumber string) error {
    r._outRoomNumber = _outRoomNumber
    r.Set("out_room_number", _outRoomNumber)
    return nil
}

// OutRoomNumber Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetOutRoomNumber() string {
    return r._outRoomNumber
}
// CheckinDate Setter
// 客人实际入住日期,opt_type=3/4时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetCheckinDate(_checkinDate string) error {
    r._checkinDate = _checkinDate
    r.Set("checkin_date", _checkinDate)
    return nil
}

// CheckinDate Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetCheckinDate() string {
    return r._checkinDate
}
// CheckoutDate Setter
// 客人实际离店日期,opt_type=4时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetCheckoutDate(_checkoutDate string) error {
    r._checkoutDate = _checkoutDate
    r.Set("checkout_date", _checkoutDate)
    return nil
}

// CheckoutDate Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetCheckoutDate() string {
    return r._checkoutDate
}
// Rooms Setter
// 客人实际预定房间数
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetRooms(_rooms int64) error {
    r._rooms = _rooms
    r.Set("rooms", _rooms)
    return nil
}

// Rooms Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetRooms() int64 {
    return r._rooms
}
// OutId Setter
// 外部订单号
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetOutId() string {
    return r._outId
}
// Tids Setter
// opt_type为10,11启用，多间房订单号列表，逗号间隔
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetTids(_tids string) error {
    r._tids = _tids
    r.Set("tids", _tids)
    return nil
}

// Tids Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetTids() string {
    return r._tids
}
// CancelType Setter
// opt_type为11启用，多间房订单取消原因类型，逗号间隔
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetCancelType(_cancelType int64) error {
    r._cancelType = _cancelType
    r.Set("cancel_type", _cancelType)
    return nil
}

// CancelType Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetCancelType() int64 {
    return r._cancelType
}
// Operator Setter
// opt_type为10,11,12启用，真实操作人
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetOperator() string {
    return r._operator
}
// ConfirmCode Setter
// opt_type为12, 订单确认号
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetConfirmCode(_confirmCode string) error {
    r._confirmCode = _confirmCode
    r.Set("confirm_code", _confirmCode)
    return nil
}

// ConfirmCode Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetConfirmCode() string {
    return r._confirmCode
}
// SelfCheckin Setter
// 是否自助入住
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetSelfCheckin(_selfCheckin bool) error {
    r._selfCheckin = _selfCheckin
    r.Set("self_checkin", _selfCheckin)
    return nil
}

// SelfCheckin Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetSelfCheckin() bool {
    return r._selfCheckin
}
// SyncToHotel Setter
// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
func (r *TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) SetSyncToHotel(_syncToHotel string) error {
    r._syncToHotel = _syncToHotel
    r.Set("sync_to_hotel", _syncToHotel)
    return nil
}

// SyncToHotel Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateAPIRequest) GetSyncToHotel() string {
    return r._syncToHotel
}
