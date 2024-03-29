package xhoteloffline

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceCheckAPIRequest 线下信用住买家资格校验接口 API请求
// taobao.xhotel.order.alipayface.check
//
// 接口用于校验买家是否具有使用酒店信用住的资格
type TaobaoXhotelOrderAlipayfaceCheckAPIRequest struct {
	model.Params
	// 证件号, 如果加密方式设置为1, 传入加密后的证件号
	_idNumber string
	// 参数必填，发布到阿里旅行的酒店编码
	_hotelCode string
	// 不清楚请留空, 用于和outHid共同定位一个酒店
	_vendor string
	// 入住人姓名
	_guestName string
	// 客人手机号
	_mobileNo string
	// 总的收费金额，单位为分
	_totalFee int64
	// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;  1: SHA-1不可逆加密,  阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串
	_encryptType int64
	// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证
	_idType int64
	// 验证类型.可以不设置. 默认0-信用住下单资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审
	_type int64
}

// NewTaobaoXhotelOrderAlipayfaceCheckRequest 初始化TaobaoXhotelOrderAlipayfaceCheckAPIRequest对象
func NewTaobaoXhotelOrderAlipayfaceCheckRequest() *TaobaoXhotelOrderAlipayfaceCheckAPIRequest {
	return &TaobaoXhotelOrderAlipayfaceCheckAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) Reset() {
	r._idNumber = ""
	r._hotelCode = ""
	r._vendor = ""
	r._guestName = ""
	r._mobileNo = ""
	r._totalFee = 0
	r._encryptType = 0
	r._idType = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdNumber is IdNumber Setter
// 证件号, 如果加密方式设置为1, 传入加密后的证件号
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetHotelCode is HotelCode Setter
// 参数必填，发布到阿里旅行的酒店编码
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetVendor is Vendor Setter
// 不清楚请留空, 用于和outHid共同定位一个酒店
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetVendor() string {
	return r._vendor
}

// SetGuestName is GuestName Setter
// 入住人姓名
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetGuestName(_guestName string) error {
	r._guestName = _guestName
	r.Set("guest_name", _guestName)
	return nil
}

// GetGuestName GuestName Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetGuestName() string {
	return r._guestName
}

// SetMobileNo is MobileNo Setter
// 客人手机号
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetMobileNo(_mobileNo string) error {
	r._mobileNo = _mobileNo
	r.Set("mobile_no", _mobileNo)
	return nil
}

// GetMobileNo MobileNo Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetMobileNo() string {
	return r._mobileNo
}

// SetTotalFee is TotalFee Setter
// 总的收费金额，单位为分
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetTotalFee(_totalFee int64) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetTotalFee() int64 {
	return r._totalFee
}

// SetEncryptType is EncryptType Setter
// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;  1: SHA-1不可逆加密,  阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetEncryptType(_encryptType int64) error {
	r._encryptType = _encryptType
	r.Set("encrypt_type", _encryptType)
	return nil
}

// GetEncryptType EncryptType Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetEncryptType() int64 {
	return r._encryptType
}

// SetIdType is IdType Setter
// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// GetIdType IdType Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetIdType() int64 {
	return r._idType
}

// SetType is Type Setter
// 验证类型.可以不设置. 默认0-信用住下单资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审
func (r *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoXhotelOrderAlipayfaceCheckAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoXhotelOrderAlipayfaceCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderAlipayfaceCheckRequest()
	},
}

// GetTaobaoXhotelOrderAlipayfaceCheckRequest 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceCheckAPIRequest
func GetTaobaoXhotelOrderAlipayfaceCheckAPIRequest() *TaobaoXhotelOrderAlipayfaceCheckAPIRequest {
	return poolTaobaoXhotelOrderAlipayfaceCheckAPIRequest.Get().(*TaobaoXhotelOrderAlipayfaceCheckAPIRequest)
}

// ReleaseTaobaoXhotelOrderAlipayfaceCheckAPIRequest 将 TaobaoXhotelOrderAlipayfaceCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceCheckAPIRequest(v *TaobaoXhotelOrderAlipayfaceCheckAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceCheckAPIRequest.Put(v)
}
