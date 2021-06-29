package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住买家资格校验接口 API请求
taobao.xhotel.order.alipayface.check

接口用于校验买家是否具有使用酒店信用住的资格
*/
type TaobaoXhotelOrderAlipayfaceCheckRequest struct {
    model.Params
    // 总的收费金额，单位为分
    totalFee   int64
    // 参数必填，发布到阿里旅行的酒店编码
    hotelCode   string
    // 证件号, 如果加密方式设置为1, 传入加密后的证件号
    idNumber   string
    // 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;  1: SHA-1不可逆加密,  阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串
    encryptType   int64
    // 验证类型.可以不设置. 默认0-信用住下单资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审
    type   int64
    // 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证
    idType   int64
    // 不清楚请留空, 用于和outHid共同定位一个酒店
    vendor   string
    // 入住人姓名
    guestName   string
    // 客人手机号
    mobileNo   string
}

// 初始化TaobaoXhotelOrderAlipayfaceCheckRequest对象
func NewTaobaoXhotelOrderAlipayfaceCheckRequest() *TaobaoXhotelOrderAlipayfaceCheckRequest{
    return &TaobaoXhotelOrderAlipayfaceCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TotalFee Setter
// 总的收费金额，单位为分
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

// TotalFee Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetTotalFee() int64 {
    return r.totalFee
}
// HotelCode Setter
// 参数必填，发布到阿里旅行的酒店编码
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetHotelCode() string {
    return r.hotelCode
}
// IdNumber Setter
// 证件号, 如果加密方式设置为1, 传入加密后的证件号
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

// IdNumber Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetIdNumber() string {
    return r.idNumber
}
// EncryptType Setter
// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;  1: SHA-1不可逆加密,  阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetEncryptType(encryptType int64) error {
    r.encryptType = encryptType
    r.Set("encrypt_type", encryptType)
    return nil
}

// EncryptType Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetEncryptType() int64 {
    return r.encryptType
}
// Type Setter
// 验证类型.可以不设置. 默认0-信用住下单资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetType() int64 {
    return r.type
}
// IdType Setter
// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

// IdType Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetIdType() int64 {
    return r.idType
}
// Vendor Setter
// 不清楚请留空, 用于和outHid共同定位一个酒店
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetVendor() string {
    return r.vendor
}
// GuestName Setter
// 入住人姓名
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetGuestName(guestName string) error {
    r.guestName = guestName
    r.Set("guest_name", guestName)
    return nil
}

// GuestName Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetGuestName() string {
    return r.guestName
}
// MobileNo Setter
// 客人手机号
func (r *TaobaoXhotelOrderAlipayfaceCheckRequest) SetMobileNo(mobileNo string) error {
    r.mobileNo = mobileNo
    r.Set("mobile_no", mobileNo)
    return nil
}

// MobileNo Getter
func (r TaobaoXhotelOrderAlipayfaceCheckRequest) GetMobileNo() string {
    return r.mobileNo
}
