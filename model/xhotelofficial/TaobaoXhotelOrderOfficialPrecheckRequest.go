package xhotelofficial

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资格预校验接口 APIRequest
taobao.xhotel.order.official.precheck

官网信用住用户资格预校验接口是在订单创建之前，根据入住人身份信息对其做预先校验是否具有信用住资格。可以优化用户预定体验，对于无资格的用户在预定前即不可进行信用住的选择。减少在提交预定后预定失败体验。该接口为可选对接接口，商家可根据实际情况自行决定是否对接。

接口使用场景

提交订单前的预定人信用住资格预先校验，卖家可决定是否在搜索，预订页，补全身份信息时进行调用，以便决定信用住是否提供给用户
*/
type TaobaoXhotelOrderOfficialPrecheckRequest struct {
    model.Params

    // 证件号, 如果加密方式设置为1, 传入加密后的证件号（建议明文传递）
    idNumber   string 

    // 总的收费金额，单位为分
    totalFee   int64 

    // 参数必填，发布到阿里旅行的酒店编码
    hotelCode   string 

    // 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息; 1: SHA-1不可逆加密, 阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串(目前不建议加密)
    encryptType   int64 

    // 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证（目前仅仅支持身份证）
    idType   int64 

    // 验证类型.可以不设置. 默认0-下单前资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审（无特殊要求不填写）
    type   int64 

    // 请咨酒店对接负责人, 用于和outHid共同定位一个酒店
    vendor   string 

    // 入住人姓名
    guestName   string 

    // 客人手机号
    mobileNo   string 

}

func NewTaobaoXhotelOrderOfficialPrecheckRequest() *TaobaoXhotelOrderOfficialPrecheckRequest{
    return &TaobaoXhotelOrderOfficialPrecheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.precheck"
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetIdNumber() string {
    return r.idNumber
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetTotalFee() int64 {
    return r.totalFee
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetEncryptType(encryptType int64) error {
    r.encryptType = encryptType
    r.Set("encrypt_type", encryptType)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetEncryptType() int64 {
    return r.encryptType
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetIdType(idType int64) error {
    r.idType = idType
    r.Set("id_type", idType)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetIdType() int64 {
    return r.idType
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetGuestName(guestName string) error {
    r.guestName = guestName
    r.Set("guest_name", guestName)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetGuestName() string {
    return r.guestName
}

func (r *TaobaoXhotelOrderOfficialPrecheckRequest) SetMobileNo(mobileNo string) error {
    r.mobileNo = mobileNo
    r.Set("mobile_no", mobileNo)
    return nil
}

func (r TaobaoXhotelOrderOfficialPrecheckRequest) GetMobileNo() string {
    return r.mobileNo
}

