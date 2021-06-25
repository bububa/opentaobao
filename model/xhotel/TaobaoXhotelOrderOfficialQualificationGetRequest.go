package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资质校验 APIRequest
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

func NewTaobaoXhotelOrderOfficialQualificationGetRequest() *TaobaoXhotelOrderOfficialQualificationGetRequest{
    return &TaobaoXhotelOrderOfficialQualificationGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.qualification.get"
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetNotifyUrl() string {
    return r.notifyUrl
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetReturnUrl(returnUrl string) error {
    r.returnUrl = returnUrl
    r.Set("return_url", returnUrl)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetReturnUrl() string {
    return r.returnUrl
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetExtendAttrs(extendAttrs string) error {
    r.extendAttrs = extendAttrs
    r.Set("extend_attrs", extendAttrs)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetExtendAttrs() string {
    return r.extendAttrs
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetMobileNo(mobileNo string) error {
    r.mobileNo = mobileNo
    r.Set("mobile_no", mobileNo)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetMobileNo() string {
    return r.mobileNo
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetIdType() int64 {
    return r.idType
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetEncryptType(encryptType int64) error {
    r.encryptType = encryptType
    r.Set("encrypt_type", encryptType)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetEncryptType() int64 {
    return r.encryptType
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetGuestName(guestName string) error {
    r.guestName = guestName
    r.Set("guest_name", guestName)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetGuestName() string {
    return r.guestName
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetAlipayAccount(alipayAccount string) error {
    r.alipayAccount = alipayAccount
    r.Set("alipay_account", alipayAccount)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetAlipayAccount() string {
    return r.alipayAccount
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetOutMemberAccount(outMemberAccount string) error {
    r.outMemberAccount = outMemberAccount
    r.Set("out_member_account", outMemberAccount)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetOutMemberAccount() string {
    return r.outMemberAccount
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetIdNumber() string {
    return r.idNumber
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetDailyPriceInfo(dailyPriceInfo string) error {
    r.dailyPriceInfo = dailyPriceInfo
    r.Set("daily_price_info", dailyPriceInfo)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetDailyPriceInfo() string {
    return r.dailyPriceInfo
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetCheckOut() string {
    return r.checkOut
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetCheckIn(checkIn string) error {
    r.checkIn = checkIn
    r.Set("check_in", checkIn)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetCheckIn() string {
    return r.checkIn
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetOutUUID(outUUID string) error {
    r.outUUID = outUUID
    r.Set("out_u_u_i_d", outUUID)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetOutUUID() string {
    return r.outUUID
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetTotalFee() int64 {
    return r.totalFee
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetOutOid(outOid string) error {
    r.outOid = outOid
    r.Set("out_oid", outOid)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetOutOid() string {
    return r.outOid
}

func (r *TaobaoXhotelOrderOfficialQualificationGetRequest) SetRoomNum(roomNum int64) error {
    r.roomNum = roomNum
    r.Set("room_num", roomNum)
    return nil
}

func (r TaobaoXhotelOrderOfficialQualificationGetRequest) GetRoomNum() int64 {
    return r.roomNum
}

