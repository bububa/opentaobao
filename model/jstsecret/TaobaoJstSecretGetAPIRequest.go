package jstsecret

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsecretgetAPIRequest 获取订单消费者的隐私号码 API请求
// taobao.jst.secret.get
//
// 根据订单号获取消费者的隐私号
type TaobaojstsecretgetAPIRequest struct {
	model.Params
	// 订单号
	_tid int64
	// 隐私号类型，1=手机号，2=分机号，默认为1；分机号使用时拨打手机号转分机号
	_type int64
	// 隐私号码过期天数，默认30天，取值范围[1,30]
	_expireDays int64
}

// NewTaobaojstsecretgetRequest 初始化TaobaojstsecretgetAPIRequest对象
func NewTaobaojstsecretgetRequest() *TaobaojstsecretgetAPIRequest {
	return &TaobaojstsecretgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsecretgetAPIRequest) GetApiMethodName() string {
	return "taobao.jst.secret.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsecretgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsecretgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 订单号
func (r *TaobaojstsecretgetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaojstsecretgetAPIRequest) GetTid() int64 {
	return r._tid
}

// SetType is Type Setter
// 隐私号类型，1=手机号，2=分机号，默认为1；分机号使用时拨打手机号转分机号
func (r *TaobaojstsecretgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaojstsecretgetAPIRequest) GetType() int64 {
	return r._type
}

// SetExpireDays is ExpireDays Setter
// 隐私号码过期天数，默认30天，取值范围[1,30]
func (r *TaobaojstsecretgetAPIRequest) SetExpireDays(_expireDays int64) error {
	r._expireDays = _expireDays
	r.Set("expire_days", _expireDays)
	return nil
}

// GetExpireDays ExpireDays Getter
func (r TaobaojstsecretgetAPIRequest) GetExpireDays() int64 {
	return r._expireDays
}
