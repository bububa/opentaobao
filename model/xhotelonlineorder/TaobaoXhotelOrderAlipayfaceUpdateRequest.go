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
type TaobaoXhotelOrderAlipayfaceUpdateRequest struct {
    model.Params
    // 淘宝订单号,必填
    tid   int64
    // 操作的类型：12.补录确认号,11.多间房确认无房，10.多间房确认有房，8.取消订单(cancel)酒店端发起取消,必须在和买家协商通过的情况下操作,否则有法务风险; 5.买家未入住(noshow),如果该单有担保，会收取买家的担保金额; 3.核实入住(checkIn); 4.核实离店(checkOut); 1.确认无房(直连卖家禁止该操作),2.确认有房(直连卖家禁止该操作)
    optType   int64
    // 无房原因分类:1.无房, 2.价格变动, 3.买家原因, 4.其它原因,opt_type=1时必填
    reasonType   int64
    // 无房原因描述:opt_type=1时必填
    reasonText   string
    // 入住房间号
    outRoomNumber   string
    // 客人实际入住日期,opt_type=3/4时必填
    checkinDate   string
    // 客人实际离店日期,opt_type=4时必填
    checkoutDate   string
    // 客人实际预定房间数
    rooms   int64
    // 外部订单号
    outId   string
    // opt_type为10,11启用，多间房订单号列表，逗号间隔
    tids   string
    // opt_type为11启用，多间房订单取消原因类型，逗号间隔
    cancelType   int64
    // opt_type为10,11,12启用，真实操作人
    operator   string
    // opt_type为12, 订单确认号
    confirmCode   string
    // 是否自助入住
    selfCheckin   bool
    // 是否把代理直签的订单同步到酒店，Y为同步，N不同步
    syncToHotel   string
}

// 初始化TaobaoXhotelOrderAlipayfaceUpdateRequest对象
func NewTaobaoXhotelOrderAlipayfaceUpdateRequest() *TaobaoXhotelOrderAlipayfaceUpdateRequest{
    return &TaobaoXhotelOrderAlipayfaceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单号,必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetTid() int64 {
    return r.tid
}
// OptType Setter
// 操作的类型：12.补录确认号,11.多间房确认无房，10.多间房确认有房，8.取消订单(cancel)酒店端发起取消,必须在和买家协商通过的情况下操作,否则有法务风险; 5.买家未入住(noshow),如果该单有担保，会收取买家的担保金额; 3.核实入住(checkIn); 4.核实离店(checkOut); 1.确认无房(直连卖家禁止该操作),2.确认有房(直连卖家禁止该操作)
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetOptType(optType int64) error {
    r.optType = optType
    r.Set("opt_type", optType)
    return nil
}

// OptType Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetOptType() int64 {
    return r.optType
}
// ReasonType Setter
// 无房原因分类:1.无房, 2.价格变动, 3.买家原因, 4.其它原因,opt_type=1时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetReasonType(reasonType int64) error {
    r.reasonType = reasonType
    r.Set("reason_type", reasonType)
    return nil
}

// ReasonType Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetReasonType() int64 {
    return r.reasonType
}
// ReasonText Setter
// 无房原因描述:opt_type=1时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetReasonText(reasonText string) error {
    r.reasonText = reasonText
    r.Set("reason_text", reasonText)
    return nil
}

// ReasonText Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetReasonText() string {
    return r.reasonText
}
// OutRoomNumber Setter
// 入住房间号
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetOutRoomNumber(outRoomNumber string) error {
    r.outRoomNumber = outRoomNumber
    r.Set("out_room_number", outRoomNumber)
    return nil
}

// OutRoomNumber Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetOutRoomNumber() string {
    return r.outRoomNumber
}
// CheckinDate Setter
// 客人实际入住日期,opt_type=3/4时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetCheckinDate(checkinDate string) error {
    r.checkinDate = checkinDate
    r.Set("checkin_date", checkinDate)
    return nil
}

// CheckinDate Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetCheckinDate() string {
    return r.checkinDate
}
// CheckoutDate Setter
// 客人实际离店日期,opt_type=4时必填
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetCheckoutDate(checkoutDate string) error {
    r.checkoutDate = checkoutDate
    r.Set("checkout_date", checkoutDate)
    return nil
}

// CheckoutDate Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetCheckoutDate() string {
    return r.checkoutDate
}
// Rooms Setter
// 客人实际预定房间数
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetRooms(rooms int64) error {
    r.rooms = rooms
    r.Set("rooms", rooms)
    return nil
}

// Rooms Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetRooms() int64 {
    return r.rooms
}
// OutId Setter
// 外部订单号
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetOutId() string {
    return r.outId
}
// Tids Setter
// opt_type为10,11启用，多间房订单号列表，逗号间隔
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetTids(tids string) error {
    r.tids = tids
    r.Set("tids", tids)
    return nil
}

// Tids Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetTids() string {
    return r.tids
}
// CancelType Setter
// opt_type为11启用，多间房订单取消原因类型，逗号间隔
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetCancelType(cancelType int64) error {
    r.cancelType = cancelType
    r.Set("cancel_type", cancelType)
    return nil
}

// CancelType Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetCancelType() int64 {
    return r.cancelType
}
// Operator Setter
// opt_type为10,11,12启用，真实操作人
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetOperator() string {
    return r.operator
}
// ConfirmCode Setter
// opt_type为12, 订单确认号
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetConfirmCode(confirmCode string) error {
    r.confirmCode = confirmCode
    r.Set("confirm_code", confirmCode)
    return nil
}

// ConfirmCode Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetConfirmCode() string {
    return r.confirmCode
}
// SelfCheckin Setter
// 是否自助入住
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetSelfCheckin(selfCheckin bool) error {
    r.selfCheckin = selfCheckin
    r.Set("self_checkin", selfCheckin)
    return nil
}

// SelfCheckin Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetSelfCheckin() bool {
    return r.selfCheckin
}
// SyncToHotel Setter
// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
func (r *TaobaoXhotelOrderAlipayfaceUpdateRequest) SetSyncToHotel(syncToHotel string) error {
    r.syncToHotel = syncToHotel
    r.Set("sync_to_hotel", syncToHotel)
    return nil
}

// SyncToHotel Getter
func (r TaobaoXhotelOrderAlipayfaceUpdateRequest) GetSyncToHotel() string {
    return r.syncToHotel
}
