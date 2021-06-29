package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资质校验 API请求
taobao.xhotel.order.official.qualification.get

官网信用住在下单前对用户进行资质校验，资质校验通过才能进行信用支付
*/
type TaobaoXhotelOrderOfficialQualificationGetRequest struct {
    model.Params
    // 卖家接收阿里旅行订单状态变更的服务地址（需要实现阿里旅行提供的服务通知规范）
    notifyUrl   string
    // 阿里旅行支付（下单）结束后跳转卖家的页面地址（必须）
    returnUrl   string
    // 扩展字段，json串，用户后续的营销、统计、定制等需求，目前已有key列表：      is_new_user：是否是卖家新用户，1-是，0或者key为null，表示不是      is_first_stay：是否是卖家首住，1-是，0或者key为null，表示不是     （已有列表必须传递）
    extendAttrs   string
    // 用户手机号(可选)
    mobileNo   string
    // 商家在淘宝给分配的渠道名（建议填充较好）
    vendor   string
    // 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证。目前只支持身份证
    idType   int64
    // 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;      * 目前只支持不加密
    encryptType   int64
    // 入住人姓名（必选）
    guestName   string
    // 用户支付宝唯一识别码(可选)
    alipayAccount   string
    // 外部会员账号（必选）
    outMemberAccount   string
    // 身份证号，必选
    idNumber   string
    // 每日房价,json格式 ，如果是多间房，则是每日多间房总房价(可选)      * eg:{"day":"2015-08-12","price":48800},      {"day":"2015-08-13","price":48800}
    dailyPriceInfo   string
    // 客人离店日期, 最多支持9间夜
    checkOut   string
    // 客人入住日期
    checkIn   string
    // 外部请求序列表号\流水号，单次请求的唯一标识(必须)
    outUUID   string
    // 总的收费金额，单位为分(必须)
    totalFee   int64
    // 酒店外部编码
    hotelCode   string
    // 外部订单号（必选），阿里旅行会根据此值进行幂等性校验
    outOid   string
    // 房间数
    roomNum   int64
}

// 初始化TaobaoXhotelOrderOfficialQualificationGetRequest对象
func NewTaobaoXhotelOrderOfficialQualificationGetRequest() *TaobaoXhotelOrderOfficialQualificationGetRequest{
    return &TaobaoXhotelOrderOfficialQualificationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.qualification.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NotifyUrl Setter
// 卖家接收阿里旅行订单状态变更的服务地址（需要实现阿里旅行提供的服务通知规范）
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetNotifyUrl() string {
    return r.notifyUrl
}
// ReturnUrl Setter
// 阿里旅行支付（下单）结束后跳转卖家的页面地址（必须）
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetReturnUrl(returnUrl string) error {
    r.returnUrl = returnUrl
    r.Set("return_url", returnUrl)
    return nil
}

// ReturnUrl Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetReturnUrl() string {
    return r.returnUrl
}
// ExtendAttrs Setter
// 扩展字段，json串，用户后续的营销、统计、定制等需求，目前已有key列表：      is_new_user：是否是卖家新用户，1-是，0或者key为null，表示不是      is_first_stay：是否是卖家首住，1-是，0或者key为null，表示不是     （已有列表必须传递）
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetExtendAttrs(extendAttrs string) error {
    r.extendAttrs = extendAttrs
    r.Set("extend_attrs", extendAttrs)
    return nil
}

// ExtendAttrs Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetExtendAttrs() string {
    return r.extendAttrs
}
// MobileNo Setter
// 用户手机号(可选)
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetMobileNo(mobileNo string) error {
    r.mobileNo = mobileNo
    r.Set("mobile_no", mobileNo)
    return nil
}

// MobileNo Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetMobileNo() string {
    return r.mobileNo
}
// Vendor Setter
// 商家在淘宝给分配的渠道名（建议填充较好）
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetVendor() string {
    return r.vendor
}
// IdType Setter
// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证。目前只支持身份证
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

// IdType Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetIdType() int64 {
    return r.idType
}
// EncryptType Setter
// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;      * 目前只支持不加密
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetEncryptType(encryptType int64) error {
    r.encryptType = encryptType
    r.Set("encrypt_type", encryptType)
    return nil
}

// EncryptType Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetEncryptType() int64 {
    return r.encryptType
}
// GuestName Setter
// 入住人姓名（必选）
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetGuestName(guestName string) error {
    r.guestName = guestName
    r.Set("guest_name", guestName)
    return nil
}

// GuestName Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetGuestName() string {
    return r.guestName
}
// AlipayAccount Setter
// 用户支付宝唯一识别码(可选)
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetAlipayAccount(alipayAccount string) error {
    r.alipayAccount = alipayAccount
    r.Set("alipay_account", alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetAlipayAccount() string {
    return r.alipayAccount
}
// OutMemberAccount Setter
// 外部会员账号（必选）
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetOutMemberAccount(outMemberAccount string) error {
    r.outMemberAccount = outMemberAccount
    r.Set("out_member_account", outMemberAccount)
    return nil
}

// OutMemberAccount Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetOutMemberAccount() string {
    return r.outMemberAccount
}
// IdNumber Setter
// 身份证号，必选
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

// IdNumber Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetIdNumber() string {
    return r.idNumber
}
// DailyPriceInfo Setter
// 每日房价,json格式 ，如果是多间房，则是每日多间房总房价(可选)      * eg:{"day":"2015-08-12","price":48800},      {"day":"2015-08-13","price":48800}
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetDailyPriceInfo(dailyPriceInfo string) error {
    r.dailyPriceInfo = dailyPriceInfo
    r.Set("daily_price_info", dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetDailyPriceInfo() string {
    return r.dailyPriceInfo
}
// CheckOut Setter
// 客人离店日期, 最多支持9间夜
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetCheckOut() string {
    return r.checkOut
}
// CheckIn Setter
// 客人入住日期
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetCheckIn(checkIn string) error {
    r.checkIn = checkIn
    r.Set("check_in", checkIn)
    return nil
}

// CheckIn Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetCheckIn() string {
    return r.checkIn
}
// OutUUID Setter
// 外部请求序列表号\流水号，单次请求的唯一标识(必须)
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetOutUUID(outUUID string) error {
    r.outUUID = outUUID
    r.Set("out_u_u_i_d", outUUID)
    return nil
}

// OutUUID Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetOutUUID() string {
    return r.outUUID
}
// TotalFee Setter
// 总的收费金额，单位为分(必须)
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

// TotalFee Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetTotalFee() int64 {
    return r.totalFee
}
// HotelCode Setter
// 酒店外部编码
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetHotelCode() string {
    return r.hotelCode
}
// OutOid Setter
// 外部订单号（必选），阿里旅行会根据此值进行幂等性校验
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetOutOid(outOid string) error {
    r.outOid = outOid
    r.Set("out_oid", outOid)
    return nil
}

// OutOid Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetOutOid() string {
    return r.outOid
}
// RoomNum Setter
// 房间数
func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetRoomNum(roomNum int64) error {
    r.roomNum = roomNum
    r.Set("room_num", roomNum)
    return nil
}

// RoomNum Getter
func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetRoomNum() int64 {
    return r.roomNum
}
