package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchantmaresendAPIRequest 电子凭证重发回调接口 API请求
// taobao.eticket.merchant.ma.resend
//
// 码商重发电子凭证回调接口
type TaobaoeticketmerchantmaresendAPIRequest struct {
	model.Params
	// 待重发的码列表
	_isvMaList []IsvMa
	// 业务id（订单号）
	_outerId string
	// 需要跟发码通知获取到的参数一致
	_token string
	// 业务类型
	_bizType int64
}

// NewTaobaoeticketmerchantmaresendRequest 初始化TaobaoeticketmerchantmaresendAPIRequest对象
func NewTaobaoeticketmerchantmaresendRequest() *TaobaoeticketmerchantmaresendAPIRequest {
	return &TaobaoeticketmerchantmaresendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoeticketmerchantmaresendAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoeticketmerchantmaresendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoeticketmerchantmaresendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvMaList is IsvMaList Setter
// 待重发的码列表
func (r *TaobaoeticketmerchantmaresendAPIRequest) SetIsvMaList(_isvMaList []IsvMa) error {
	r._isvMaList = _isvMaList
	r.Set("isv_ma_list", _isvMaList)
	return nil
}

// GetIsvMaList IsvMaList Getter
func (r TaobaoeticketmerchantmaresendAPIRequest) GetIsvMaList() []IsvMa {
	return r._isvMaList
}

// SetOuterId is OuterId Setter
// 业务id（订单号）
func (r *TaobaoeticketmerchantmaresendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoeticketmerchantmaresendAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetToken is Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoeticketmerchantmaresendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoeticketmerchantmaresendAPIRequest) GetToken() string {
	return r._token
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoeticketmerchantmaresendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoeticketmerchantmaresendAPIRequest) GetBizType() int64 {
	return r._bizType
}
