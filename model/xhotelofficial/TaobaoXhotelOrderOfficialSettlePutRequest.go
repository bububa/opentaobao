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
type TaobaoXhotelOrderOfficialSettlePutRequest struct {
    model.Params
    // 淘宝订单id,必须填写
    tid   int64
    // 房费总额(必须大于0)
    totalRoomFee   int64
    // 杂费总额(不能为负数)
    otherFee   int64
    // 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
    otherFeeDetail   string
    // 商家订单号
    outId   string
    // 入住房间号
    roomNo   string
    // 每日房价,json格式,如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
    dailyPriceInfo   string
    // 实际离店日期，用于校验总房费收取
    checkOut   string
    // 备注
    memo   string
    // 房间明细列表
    roomSettleInfoList   []RoomSettleInfo
    // 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含（多间房结账必须传入）
    containGuarantee   int64
    // 结账请求流水号
    outUuid   string
    // 请求结果通知地址（暂时无效，无需传入）
    notifyUrl   string
}

// 初始化TaobaoXhotelOrderOfficialSettlePutRequest对象
func NewTaobaoXhotelOrderOfficialSettlePutRequest() *TaobaoXhotelOrderOfficialSettlePutRequest{
    return &TaobaoXhotelOrderOfficialSettlePutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.settle.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单id,必须填写
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetTid() int64 {
    return r.tid
}
// TotalRoomFee Setter
// 房费总额(必须大于0)
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetTotalRoomFee(totalRoomFee int64) error {
    r.totalRoomFee = totalRoomFee
    r.Set("total_room_fee", totalRoomFee)
    return nil
}

// TotalRoomFee Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetTotalRoomFee() int64 {
    return r.totalRoomFee
}
// OtherFee Setter
// 杂费总额(不能为负数)
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetOtherFee(otherFee int64) error {
    r.otherFee = otherFee
    r.Set("other_fee", otherFee)
    return nil
}

// OtherFee Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetOtherFee() int64 {
    return r.otherFee
}
// OtherFeeDetail Setter
// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetOtherFeeDetail(otherFeeDetail string) error {
    r.otherFeeDetail = otherFeeDetail
    r.Set("other_fee_detail", otherFeeDetail)
    return nil
}

// OtherFeeDetail Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetOtherFeeDetail() string {
    return r.otherFeeDetail
}
// OutId Setter
// 商家订单号
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetOutId() string {
    return r.outId
}
// RoomNo Setter
// 入住房间号
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetRoomNo(roomNo string) error {
    r.roomNo = roomNo
    r.Set("room_no", roomNo)
    return nil
}

// RoomNo Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetRoomNo() string {
    return r.roomNo
}
// DailyPriceInfo Setter
// 每日房价,json格式,如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetDailyPriceInfo(dailyPriceInfo string) error {
    r.dailyPriceInfo = dailyPriceInfo
    r.Set("daily_price_info", dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetDailyPriceInfo() string {
    return r.dailyPriceInfo
}
// CheckOut Setter
// 实际离店日期，用于校验总房费收取
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetCheckOut() string {
    return r.checkOut
}
// Memo Setter
// 备注
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetMemo() string {
    return r.memo
}
// RoomSettleInfoList Setter
// 房间明细列表
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetRoomSettleInfoList(roomSettleInfoList []RoomSettleInfo) error {
    r.roomSettleInfoList = roomSettleInfoList
    r.Set("room_settle_info_list", roomSettleInfoList)
    return nil
}

// RoomSettleInfoList Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetRoomSettleInfoList() []RoomSettleInfo {
    return r.roomSettleInfoList
}
// ContainGuarantee Setter
// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含（多间房结账必须传入）
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetContainGuarantee(containGuarantee int64) error {
    r.containGuarantee = containGuarantee
    r.Set("contain_guarantee", containGuarantee)
    return nil
}

// ContainGuarantee Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetContainGuarantee() int64 {
    return r.containGuarantee
}
// OutUuid Setter
// 结账请求流水号
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetOutUuid() string {
    return r.outUuid
}
// NotifyUrl Setter
// 请求结果通知地址（暂时无效，无需传入）
func (r *TaobaoXhotelOrderOfficialSettlePutRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialSettlePutRequest) GetNotifyUrl() string {
    return r.notifyUrl
}
