package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderofficialqualificationgetAPIRequest 官网信用住用户资质校验 API请求
// taobao.xhotel.order.official.qualification.get
//
// 官网信用住在下单前对用户进行资质校验，资质校验通过才能进行信用支付
type TaobaoxhotelorderofficialqualificationgetAPIRequest struct {
	model.Params
	// 外部请求序列表号\流水号，单次请求的唯一标识(必须)
	_outUUID string
	// 酒店外部编码
	_hotelCode string
	// 客人离店日期, 最多支持9间夜
	_checkOut string
	// 身份证号，必选
	_idNumber string
	// 每日房价,json格式 ，如果是多间房，则是每日多间房总房价(可选)      * eg:{"day":"2015-08-12","price":48800},      {"day":"2015-08-13","price":48800}
	_dailyPriceInfo string
	// 外部会员账号（必选）
	_outMemberAccount string
	// 用户支付宝唯一识别码(可选)
	_alipayAccount string
	// 入住人姓名（必选）
	_guestName string
	// 商家在淘宝给分配的渠道名（建议填充较好）
	_vendor string
	// 用户手机号(可选)
	_mobileNo string
	// 扩展字段，json串，用户后续的营销、统计、定制等需求，目前已有key列表：      is_new_user：是否是卖家新用户，1-是，0或者key为null，表示不是      is_first_stay：是否是卖家首住，1-是，0或者key为null，表示不是     （已有列表必须传递）
	_extendAttrs string
	// 阿里旅行支付（下单）结束后跳转卖家的页面地址（必须）
	_returnUrl string
	// 卖家接收阿里旅行订单状态变更的服务地址（需要实现阿里旅行提供的服务通知规范）
	_notifyUrl string
	// 客人入住日期
	_checkIn string
	// 外部订单号（必选），阿里旅行会根据此值进行幂等性校验
	_outOid string
	// 总的收费金额，单位为分(必须)
	_totalFee int64
	// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;      * 目前只支持不加密
	_encryptType int64
	// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证。目前只支持身份证
	_idType int64
	// 房间数
	_roomNum int64
}

// NewTaobaoxhotelorderofficialqualificationgetRequest 初始化TaobaoxhotelorderofficialqualificationgetAPIRequest对象
func NewTaobaoxhotelorderofficialqualificationgetRequest() *TaobaoxhotelorderofficialqualificationgetAPIRequest {
	return &TaobaoxhotelorderofficialqualificationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.official.qualification.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutUUID is OutUUID Setter
// 外部请求序列表号\流水号，单次请求的唯一标识(必须)
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetOutUUID(_outUUID string) error {
	r._outUUID = _outUUID
	r.Set("out_u_u_i_d", _outUUID)
	return nil
}

// GetOutUUID OutUUID Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetOutUUID() string {
	return r._outUUID
}

// SetHotelCode is HotelCode Setter
// 酒店外部编码
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetCheckOut is CheckOut Setter
// 客人离店日期, 最多支持9间夜
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetIdNumber is IdNumber Setter
// 身份证号，必选
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetDailyPriceInfo is DailyPriceInfo Setter
// 每日房价,json格式 ，如果是多间房，则是每日多间房总房价(可选)      * eg:{&#34;day&#34;:&#34;2015-08-12&#34;,&#34;price&#34;:48800},      {&#34;day&#34;:&#34;2015-08-13&#34;,&#34;price&#34;:48800}
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
	r._dailyPriceInfo = _dailyPriceInfo
	r.Set("daily_price_info", _dailyPriceInfo)
	return nil
}

// GetDailyPriceInfo DailyPriceInfo Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetDailyPriceInfo() string {
	return r._dailyPriceInfo
}

// SetOutMemberAccount is OutMemberAccount Setter
// 外部会员账号（必选）
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetOutMemberAccount(_outMemberAccount string) error {
	r._outMemberAccount = _outMemberAccount
	r.Set("out_member_account", _outMemberAccount)
	return nil
}

// GetOutMemberAccount OutMemberAccount Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetOutMemberAccount() string {
	return r._outMemberAccount
}

// SetAlipayAccount is AlipayAccount Setter
// 用户支付宝唯一识别码(可选)
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetAlipayAccount(_alipayAccount string) error {
	r._alipayAccount = _alipayAccount
	r.Set("alipay_account", _alipayAccount)
	return nil
}

// GetAlipayAccount AlipayAccount Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetAlipayAccount() string {
	return r._alipayAccount
}

// SetGuestName is GuestName Setter
// 入住人姓名（必选）
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetGuestName(_guestName string) error {
	r._guestName = _guestName
	r.Set("guest_name", _guestName)
	return nil
}

// GetGuestName GuestName Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetGuestName() string {
	return r._guestName
}

// SetVendor is Vendor Setter
// 商家在淘宝给分配的渠道名（建议填充较好）
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetMobileNo is MobileNo Setter
// 用户手机号(可选)
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetMobileNo(_mobileNo string) error {
	r._mobileNo = _mobileNo
	r.Set("mobile_no", _mobileNo)
	return nil
}

// GetMobileNo MobileNo Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetMobileNo() string {
	return r._mobileNo
}

// SetExtendAttrs is ExtendAttrs Setter
// 扩展字段，json串，用户后续的营销、统计、定制等需求，目前已有key列表：      is_new_user：是否是卖家新用户，1-是，0或者key为null，表示不是      is_first_stay：是否是卖家首住，1-是，0或者key为null，表示不是     （已有列表必须传递）
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetExtendAttrs(_extendAttrs string) error {
	r._extendAttrs = _extendAttrs
	r.Set("extend_attrs", _extendAttrs)
	return nil
}

// GetExtendAttrs ExtendAttrs Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetExtendAttrs() string {
	return r._extendAttrs
}

// SetReturnUrl is ReturnUrl Setter
// 阿里旅行支付（下单）结束后跳转卖家的页面地址（必须）
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetReturnUrl(_returnUrl string) error {
	r._returnUrl = _returnUrl
	r.Set("return_url", _returnUrl)
	return nil
}

// GetReturnUrl ReturnUrl Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetReturnUrl() string {
	return r._returnUrl
}

// SetNotifyUrl is NotifyUrl Setter
// 卖家接收阿里旅行订单状态变更的服务地址（需要实现阿里旅行提供的服务通知规范）
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetNotifyUrl(_notifyUrl string) error {
	r._notifyUrl = _notifyUrl
	r.Set("notify_url", _notifyUrl)
	return nil
}

// GetNotifyUrl NotifyUrl Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetNotifyUrl() string {
	return r._notifyUrl
}

// SetCheckIn is CheckIn Setter
// 客人入住日期
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetOutOid is OutOid Setter
// 外部订单号（必选），阿里旅行会根据此值进行幂等性校验
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetOutOid(_outOid string) error {
	r._outOid = _outOid
	r.Set("out_oid", _outOid)
	return nil
}

// GetOutOid OutOid Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetOutOid() string {
	return r._outOid
}

// SetTotalFee is TotalFee Setter
// 总的收费金额，单位为分(必须)
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetTotalFee(_totalFee int64) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetTotalFee() int64 {
	return r._totalFee
}

// SetEncryptType is EncryptType Setter
// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;      * 目前只支持不加密
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetEncryptType(_encryptType int64) error {
	r._encryptType = _encryptType
	r.Set("encrypt_type", _encryptType)
	return nil
}

// GetEncryptType EncryptType Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetEncryptType() int64 {
	return r._encryptType
}

// SetIdType is IdType Setter
// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证。目前只支持身份证
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetIdType(_idType int64) error {
	r._idType = _idType
	r.Set("id_type", _idType)
	return nil
}

// GetIdType IdType Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetIdType() int64 {
	return r._idType
}

// SetRoomNum is RoomNum Setter
// 房间数
func (r *TaobaoxhotelorderofficialqualificationgetAPIRequest) SetRoomNum(_roomNum int64) error {
	r._roomNum = _roomNum
	r.Set("room_num", _roomNum)
	return nil
}

// GetRoomNum RoomNum Getter
func (r TaobaoxhotelorderofficialqualificationgetAPIRequest) GetRoomNum() int64 {
	return r._roomNum
}
