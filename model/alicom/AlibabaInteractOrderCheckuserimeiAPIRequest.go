package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractOrderCheckuserimeiAPIRequest 金融购机验证设备号 API请求
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
type AlibabaInteractOrderCheckuserimeiAPIRequest struct {
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

// NewAlibabaInteractOrderCheckuserimeiRequest 初始化AlibabaInteractOrderCheckuserimeiAPIRequest对象
func NewAlibabaInteractOrderCheckuserimeiRequest() *AlibabaInteractOrderCheckuserimeiAPIRequest {
	return &AlibabaInteractOrderCheckuserimeiAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) Reset() {
	r._userNick = ""
	r._name = ""
	r._cardNo = ""
	r._imeis = ""
	r._mtopImei = ""
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.order.checkuserimei"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 淘宝NICK
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetName is Name Setter
// 姓名
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetName() string {
	return r._name
}

// SetCardNo is CardNo Setter
// 证件号码
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetCardNo() string {
	return r._cardNo
}

// SetImeis is Imeis Setter
// 商家传入的设备号
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) SetImeis(_imeis string) error {
	r._imeis = _imeis
	r.Set("imeis", _imeis)
	return nil
}

// GetImeis Imeis Getter
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetImeis() string {
	return r._imeis
}

// SetMtopImei is MtopImei Setter
// mtop获取的用户设备号
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) SetMtopImei(_mtopImei string) error {
	r._mtopImei = _mtopImei
	r.Set("mtop_imei", _mtopImei)
	return nil
}

// GetMtopImei MtopImei Getter
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetMtopImei() string {
	return r._mtopImei
}

// SetUserId is UserId Setter
// 淘宝UID
func (r *AlibabaInteractOrderCheckuserimeiAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaInteractOrderCheckuserimeiAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlibabaInteractOrderCheckuserimeiAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractOrderCheckuserimeiRequest()
	},
}

// GetAlibabaInteractOrderCheckuserimeiRequest 从 sync.Pool 获取 AlibabaInteractOrderCheckuserimeiAPIRequest
func GetAlibabaInteractOrderCheckuserimeiAPIRequest() *AlibabaInteractOrderCheckuserimeiAPIRequest {
	return poolAlibabaInteractOrderCheckuserimeiAPIRequest.Get().(*AlibabaInteractOrderCheckuserimeiAPIRequest)
}

// ReleaseAlibabaInteractOrderCheckuserimeiAPIRequest 将 AlibabaInteractOrderCheckuserimeiAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractOrderCheckuserimeiAPIRequest(v *AlibabaInteractOrderCheckuserimeiAPIRequest) {
	v.Reset()
	poolAlibabaInteractOrderCheckuserimeiAPIRequest.Put(v)
}
