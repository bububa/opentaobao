package xhotelofficial

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资格预校验接口 API请求
taobao.xhotel.order.official.precheck

官网信用住用户资格预校验接口是在订单创建之前，根据入住人身份信息对其做预先校验是否具有信用住资格。可以优化用户预定体验，对于无资格的用户在预定前即不可进行信用住的选择。减少在提交预定后预定失败体验。该接口为可选对接接口，商家可根据实际情况自行决定是否对接。

接口使用场景

提交订单前的预定人信用住资格预先校验，卖家可决定是否在搜索，预订页，补全身份信息时进行调用，以便决定信用住是否提供给用户
*/
type TaobaoXhotelOrderOfficialPrecheckRequest struct {
    model.Params
    // 证件号, 如果加密方式设置为1, 传入加密后的证件号（建议明文传递）
    _idNumber   string
    // 总的收费金额，单位为分
    _totalFee   int64
    // 参数必填，发布到阿里旅行的酒店编码
    _hotelCode   string
    // 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息; 1: SHA-1不可逆加密, 阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串(目前不建议加密)
    _encryptType   int64
    // 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证（目前仅仅支持身份证）
    _idType   int64
    // 验证类型.可以不设置. 默认0-下单前资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审（无特殊要求不填写）
    _type   int64
    // 请咨酒店对接负责人, 用于和outHid共同定位一个酒店
    _vendor   string
    // 入住人姓名
    _guestName   string
    // 客人手机号
    _mobileNo   string
}

// 初始化TaobaoXhotelOrderOfficialPrecheckRequest对象
func NewTaobaoXhotelOrderOfficialPrecheckRequest() *TaobaoXhotelOrderOfficialPrecheckRequest{
    return &TaobaoXhotelOrderOfficialPrecheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.precheck"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdNumber Setter
// 证件号, 如果加密方式设置为1, 传入加密后的证件号（建议明文传递）
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetIdNumber(_idNumber string) error {
    r._idNumber = _idNumber
    r.Set("id_number", _idNumber)
    return nil
}

// IdNumber Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetIdNumber() string {
    return r._idNumber
}
// TotalFee Setter
// 总的收费金额，单位为分
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetTotalFee(_totalFee int64) error {
    r._totalFee = _totalFee
    r.Set("total_fee", _totalFee)
    return nil
}

// TotalFee Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetTotalFee() int64 {
    return r._totalFee
}
// HotelCode Setter
// 参数必填，发布到阿里旅行的酒店编码
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetHotelCode() string {
    return r._hotelCode
}
// EncryptType Setter
// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息; 1: SHA-1不可逆加密, 阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串(目前不建议加密)
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetEncryptType(_encryptType int64) error {
    r._encryptType = _encryptType
    r.Set("encrypt_type", _encryptType)
    return nil
}

// EncryptType Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetEncryptType() int64 {
    return r._encryptType
}
// IdType Setter
// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证（目前仅仅支持身份证）
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetIdType(_idType int64) error {
    r._idType = _idType
    r.Set("id_type", _idType)
    return nil
}

// IdType Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetIdType() int64 {
    return r._idType
}
// Type Setter
// 验证类型.可以不设置. 默认0-下单前资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审（无特殊要求不填写）
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetType() int64 {
    return r._type
}
// Vendor Setter
// 请咨酒店对接负责人, 用于和outHid共同定位一个酒店
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetVendor() string {
    return r._vendor
}
// GuestName Setter
// 入住人姓名
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetGuestName(_guestName string) error {
    r._guestName = _guestName
    r.Set("guest_name", _guestName)
    return nil
}

// GuestName Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetGuestName() string {
    return r._guestName
}
// MobileNo Setter
// 客人手机号
func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetMobileNo(_mobileNo string) error {
    r._mobileNo = _mobileNo
    r.Set("mobile_no", _mobileNo)
    return nil
}

// MobileNo Getter
func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetMobileNo() string {
    return r._mobileNo
}
