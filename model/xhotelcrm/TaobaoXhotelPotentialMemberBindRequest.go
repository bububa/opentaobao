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
    firstName   string
    // 姓
    lastName   string
    // 电话
    phone   string
    // 邮箱
    email   string
    // 卡号
    cardNo   string
    // 等级(V1,V2,V3)
    grade   string
    // 注册时间
    registerDate   string
    // 生效时间
    fromDate   string
    // 截止时间
    toDate   string
    // 性别(M,F,U-未知)
    sex   string
    // 城市
    city   string
    // 年龄
    age   int64
    // 籍贯
    nativePlace   string
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
func (r *TaobaoXhotelPotentialMemberBindRequest) SetFirstName(firstName string) error {
    r.firstName = firstName
    r.Set("first_name", firstName)
    return nil
}

// FirstName Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetFirstName() string {
    return r.firstName
}
// LastName Setter
// 姓
func (r *TaobaoXhotelPotentialMemberBindRequest) SetLastName(lastName string) error {
    r.lastName = lastName
    r.Set("last_name", lastName)
    return nil
}

// LastName Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetLastName() string {
    return r.lastName
}
// Phone Setter
// 电话
func (r *TaobaoXhotelPotentialMemberBindRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetPhone() string {
    return r.phone
}
// Email Setter
// 邮箱
func (r *TaobaoXhotelPotentialMemberBindRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

// Email Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetEmail() string {
    return r.email
}
// CardNo Setter
// 卡号
func (r *TaobaoXhotelPotentialMemberBindRequest) SetCardNo(cardNo string) error {
    r.cardNo = cardNo
    r.Set("card_no", cardNo)
    return nil
}

// CardNo Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetCardNo() string {
    return r.cardNo
}
// Grade Setter
// 等级(V1,V2,V3)
func (r *TaobaoXhotelPotentialMemberBindRequest) SetGrade(grade string) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

// Grade Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetGrade() string {
    return r.grade
}
// RegisterDate Setter
// 注册时间
func (r *TaobaoXhotelPotentialMemberBindRequest) SetRegisterDate(registerDate string) error {
    r.registerDate = registerDate
    r.Set("register_date", registerDate)
    return nil
}

// RegisterDate Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetRegisterDate() string {
    return r.registerDate
}
// FromDate Setter
// 生效时间
func (r *TaobaoXhotelPotentialMemberBindRequest) SetFromDate(fromDate string) error {
    r.fromDate = fromDate
    r.Set("from_date", fromDate)
    return nil
}

// FromDate Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetFromDate() string {
    return r.fromDate
}
// ToDate Setter
// 截止时间
func (r *TaobaoXhotelPotentialMemberBindRequest) SetToDate(toDate string) error {
    r.toDate = toDate
    r.Set("to_date", toDate)
    return nil
}

// ToDate Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetToDate() string {
    return r.toDate
}
// Sex Setter
// 性别(M,F,U-未知)
func (r *TaobaoXhotelPotentialMemberBindRequest) SetSex(sex string) error {
    r.sex = sex
    r.Set("sex", sex)
    return nil
}

// Sex Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetSex() string {
    return r.sex
}
// City Setter
// 城市
func (r *TaobaoXhotelPotentialMemberBindRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetCity() string {
    return r.city
}
// Age Setter
// 年龄
func (r *TaobaoXhotelPotentialMemberBindRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

// Age Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetAge() int64 {
    return r.age
}
// NativePlace Setter
// 籍贯
func (r *TaobaoXhotelPotentialMemberBindRequest) SetNativePlace(nativePlace string) error {
    r.nativePlace = nativePlace
    r.Set("native_place", nativePlace)
    return nil
}

// NativePlace Getter
func (r TaobaoXhotelPotentialMemberBindRequest) GetNativePlace() string {
    return r.nativePlace
}
