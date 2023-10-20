package xhotelcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelpotentialmemberbindAPIRequest 飞猪酒店商家会员绑定 API请求
// taobao.xhotel.potential.member.bind
//
// 支持互通商家发起会员绑定
type TaobaoxhotelpotentialmemberbindAPIRequest struct {
	model.Params
	// 名
	_firstName string
	// 姓
	_lastName string
	// 电话
	_phone string
	// 邮箱
	_email string
	// 卡号
	_cardNo string
	// 等级(V1,V2,V3)
	_grade string
	// 注册时间
	_registerDate string
	// 生效时间
	_fromDate string
	// 截止时间
	_toDate string
	// 性别(M,F,U-未知)
	_sex string
	// 城市
	_city string
	// 籍贯
	_nativePlace string
	// 年龄
	_age int64
}

// NewTaobaoxhotelpotentialmemberbindRequest 初始化TaobaoxhotelpotentialmemberbindAPIRequest对象
func NewTaobaoxhotelpotentialmemberbindRequest() *TaobaoxhotelpotentialmemberbindAPIRequest {
	return &TaobaoxhotelpotentialmemberbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.potential.member.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFirstName is FirstName Setter
// 名
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetFirstName(_firstName string) error {
	r._firstName = _firstName
	r.Set("first_name", _firstName)
	return nil
}

// GetFirstName FirstName Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetFirstName() string {
	return r._firstName
}

// SetLastName is LastName Setter
// 姓
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetLastName(_lastName string) error {
	r._lastName = _lastName
	r.Set("last_name", _lastName)
	return nil
}

// GetLastName LastName Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetLastName() string {
	return r._lastName
}

// SetPhone is Phone Setter
// 电话
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetPhone() string {
	return r._phone
}

// SetEmail is Email Setter
// 邮箱
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetEmail() string {
	return r._email
}

// SetCardNo is CardNo Setter
// 卡号
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetCardNo() string {
	return r._cardNo
}

// SetGrade is Grade Setter
// 等级(V1,V2,V3)
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetGrade(_grade string) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetGrade() string {
	return r._grade
}

// SetRegisterDate is RegisterDate Setter
// 注册时间
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetRegisterDate(_registerDate string) error {
	r._registerDate = _registerDate
	r.Set("register_date", _registerDate)
	return nil
}

// GetRegisterDate RegisterDate Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetRegisterDate() string {
	return r._registerDate
}

// SetFromDate is FromDate Setter
// 生效时间
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetFromDate(_fromDate string) error {
	r._fromDate = _fromDate
	r.Set("from_date", _fromDate)
	return nil
}

// GetFromDate FromDate Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetFromDate() string {
	return r._fromDate
}

// SetToDate is ToDate Setter
// 截止时间
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetToDate(_toDate string) error {
	r._toDate = _toDate
	r.Set("to_date", _toDate)
	return nil
}

// GetToDate ToDate Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetToDate() string {
	return r._toDate
}

// SetSex is Sex Setter
// 性别(M,F,U-未知)
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetSex(_sex string) error {
	r._sex = _sex
	r.Set("sex", _sex)
	return nil
}

// GetSex Sex Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetSex() string {
	return r._sex
}

// SetCity is City Setter
// 城市
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetCity() string {
	return r._city
}

// SetNativePlace is NativePlace Setter
// 籍贯
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetNativePlace(_nativePlace string) error {
	r._nativePlace = _nativePlace
	r.Set("native_place", _nativePlace)
	return nil
}

// GetNativePlace NativePlace Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetNativePlace() string {
	return r._nativePlace
}

// SetAge is Age Setter
// 年龄
func (r *TaobaoxhotelpotentialmemberbindAPIRequest) SetAge(_age int64) error {
	r._age = _age
	r.Set("age", _age)
	return nil
}

// GetAge Age Getter
func (r TaobaoxhotelpotentialmemberbindAPIRequest) GetAge() int64 {
	return r._age
}
