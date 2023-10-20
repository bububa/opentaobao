package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractordercheckuserimeiAPIRequest 金融购机验证设备号 API请求
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
type AlibabainteractordercheckuserimeiAPIRequest struct {
	model.Params
	// 淘宝NICK
	_userNick string
	// 姓名
	_name string
	// 证件号码
	_cardNo string
	// 商家传入的设备号
	_imeis string
	// mtop获取的用户设备号
	_mtopImei string
	// 淘宝UID
	_userId int64
}

// NewAlibabainteractordercheckuserimeiRequest 初始化AlibabainteractordercheckuserimeiAPIRequest对象
func NewAlibabainteractordercheckuserimeiRequest() *AlibabainteractordercheckuserimeiAPIRequest {
	return &AlibabainteractordercheckuserimeiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractordercheckuserimeiAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.order.checkuserimei"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractordercheckuserimeiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractordercheckuserimeiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 淘宝NICK
func (r *AlibabainteractordercheckuserimeiAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabainteractordercheckuserimeiAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetName is Name Setter
// 姓名
func (r *AlibabainteractordercheckuserimeiAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabainteractordercheckuserimeiAPIRequest) GetName() string {
	return r._name
}

// SetCardNo is CardNo Setter
// 证件号码
func (r *AlibabainteractordercheckuserimeiAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r AlibabainteractordercheckuserimeiAPIRequest) GetCardNo() string {
	return r._cardNo
}

// SetImeis is Imeis Setter
// 商家传入的设备号
func (r *AlibabainteractordercheckuserimeiAPIRequest) SetImeis(_imeis string) error {
	r._imeis = _imeis
	r.Set("imeis", _imeis)
	return nil
}

// GetImeis Imeis Getter
func (r AlibabainteractordercheckuserimeiAPIRequest) GetImeis() string {
	return r._imeis
}

// SetMtopImei is MtopImei Setter
// mtop获取的用户设备号
func (r *AlibabainteractordercheckuserimeiAPIRequest) SetMtopImei(_mtopImei string) error {
	r._mtopImei = _mtopImei
	r.Set("mtop_imei", _mtopImei)
	return nil
}

// GetMtopImei MtopImei Getter
func (r AlibabainteractordercheckuserimeiAPIRequest) GetMtopImei() string {
	return r._mtopImei
}

// SetUserId is UserId Setter
// 淘宝UID
func (r *AlibabainteractordercheckuserimeiAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabainteractordercheckuserimeiAPIRequest) GetUserId() int64 {
	return r._userId
}
