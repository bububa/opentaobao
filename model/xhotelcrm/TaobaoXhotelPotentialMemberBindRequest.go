package xhotelcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪酒店商家会员绑定 API请求
taobao.xhotel.potential.member.bind

支持互通商家发起会员绑定
*/
type TaobaoXhotelPotentialMemberBindRequest struct {
    model.Params
    // 名
    _firstName   string
    // 姓
    _lastName   string
    // 电话
    _phone   string
    // 邮箱
    _email   string
    // 卡号
    _cardNo   string
    // 等级(V1,V2,V3)
    _grade   string
    // 注册时间
    _registerDate   string
    // 生效时间
    _fromDate   string
    // 截止时间
    _toDate   string
    // 性别(M,F,U-未知)
    _sex   string
    // 城市
    _city   string
    // 年龄
    _age   int64
    // 籍贯
    _nativePlace   string
}

// 初始化TaobaoXhotelPotentialMemberBindRequest对象
func NewTaobaoXhotelPotentialMemberBindRequest() *TaobaoXhotelPotentialMemberBindRequest{
    return &TaobaoXhotelPotentialMemberBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelPotentialMemberBindRequest) GetApiMethodName() string {
    return "taobao.xhotel.potential.member.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelPotentialMemberBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FirstName Setter
// 名
func (r *TaobaoXhotelPotentialMemberBindRequest) SetFirstName(_firstName string) error {
    r._firstName = _firstName
    r.Set("first_name", _firstName)
    return nil
}

// FirstName Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetFirstName() string {
    return r._firstName
}
// LastName Setter
// 姓
func (r *TaobaoXhotelPotentialMemberBindRequest) SetLastName(_lastName string) error {
    r._lastName = _lastName
    r.Set("last_name", _lastName)
    return nil
}

// LastName Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetLastName() string {
    return r._lastName
}
// Phone Setter
// 电话
func (r *TaobaoXhotelPotentialMemberBindRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetPhone() string {
    return r._phone
}
// Email Setter
// 邮箱
func (r *TaobaoXhotelPotentialMemberBindRequest) SetEmail(_email string) error {
    r._email = _email
    r.Set("email", _email)
    return nil
}

// Email Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetEmail() string {
    return r._email
}
// CardNo Setter
// 卡号
func (r *TaobaoXhotelPotentialMemberBindRequest) SetCardNo(_cardNo string) error {
    r._cardNo = _cardNo
    r.Set("card_no", _cardNo)
    return nil
}

// CardNo Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetCardNo() string {
    return r._cardNo
}
// Grade Setter
// 等级(V1,V2,V3)
func (r *TaobaoXhotelPotentialMemberBindRequest) SetGrade(_grade string) error {
    r._grade = _grade
    r.Set("grade", _grade)
    return nil
}

// Grade Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetGrade() string {
    return r._grade
}
// RegisterDate Setter
// 注册时间
func (r *TaobaoXhotelPotentialMemberBindRequest) SetRegisterDate(_registerDate string) error {
    r._registerDate = _registerDate
    r.Set("register_date", _registerDate)
    return nil
}

// RegisterDate Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetRegisterDate() string {
    return r._registerDate
}
// FromDate Setter
// 生效时间
func (r *TaobaoXhotelPotentialMemberBindRequest) SetFromDate(_fromDate string) error {
    r._fromDate = _fromDate
    r.Set("from_date", _fromDate)
    return nil
}

// FromDate Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetFromDate() string {
    return r._fromDate
}
// ToDate Setter
// 截止时间
func (r *TaobaoXhotelPotentialMemberBindRequest) SetToDate(_toDate string) error {
    r._toDate = _toDate
    r.Set("to_date", _toDate)
    return nil
}

// ToDate Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetToDate() string {
    return r._toDate
}
// Sex Setter
// 性别(M,F,U-未知)
func (r *TaobaoXhotelPotentialMemberBindRequest) SetSex(_sex string) error {
    r._sex = _sex
    r.Set("sex", _sex)
    return nil
}

// Sex Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetSex() string {
    return r._sex
}
// City Setter
// 城市
func (r *TaobaoXhotelPotentialMemberBindRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetCity() string {
    return r._city
}
// Age Setter
// 年龄
func (r *TaobaoXhotelPotentialMemberBindRequest) SetAge(_age int64) error {
    r._age = _age
    r.Set("age", _age)
    return nil
}

// Age Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetAge() int64 {
    return r._age
}
// NativePlace Setter
// 籍贯
func (r *TaobaoXhotelPotentialMemberBindRequest) SetNativePlace(_nativePlace string) error {
    r._nativePlace = _nativePlace
    r.Set("native_place", _nativePlace)
    return nil
}

// NativePlace Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetNativePlace() string {
    return r._nativePlace
}
