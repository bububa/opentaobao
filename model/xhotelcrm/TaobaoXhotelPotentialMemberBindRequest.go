package xhotelcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪酒店商家会员绑定 APIRequest
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

func NewTaobaoXhotelPotentialMemberBindRequest() *TaobaoXhotelPotentialMemberBindRequest{
    return &TaobaoXhotelPotentialMemberBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetApiMethodName() string {
    return "taobao.xhotel.potential.member.bind"
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelPotentialMemberBindRequest) SetFirstName(firstName string) error {
    r.firstName = firstName
    r.Set("first_name", firstName)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetFirstName() string {
    return r.firstName
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetLastName(lastName string) error {
    r.lastName = lastName
    r.Set("last_name", lastName)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetLastName() string {
    return r.lastName
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetPhone() string {
    return r.phone
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetEmail() string {
    return r.email
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetCardNo(cardNo string) error {
    r.cardNo = cardNo
    r.Set("card_no", cardNo)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetCardNo() string {
    return r.cardNo
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetGrade(grade string) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetGrade() string {
    return r.grade
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetRegisterDate(registerDate string) error {
    r.registerDate = registerDate
    r.Set("register_date", registerDate)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetRegisterDate() string {
    return r.registerDate
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetFromDate(fromDate string) error {
    r.fromDate = fromDate
    r.Set("from_date", fromDate)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetFromDate() string {
    return r.fromDate
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetToDate(toDate string) error {
    r.toDate = toDate
    r.Set("to_date", toDate)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetToDate() string {
    return r.toDate
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetSex(sex string) error {
    r.sex = sex
    r.Set("sex", sex)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetSex() string {
    return r.sex
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetCity() string {
    return r.city
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetAge(age int64) error {
    r.age = age
    r.Set("age", age)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetAge() int64 {
    return r.age
}

func (r *TaobaoXhotelPotentialMemberBindRequest) SetNativePlace(nativePlace string) error {
    r.nativePlace = nativePlace
    r.Set("native_place", nativePlace)
    return nil
}

func (r TaobaoXhotelPotentialMemberBindRequest) GetNativePlace() string {
    return r.nativePlace
}

