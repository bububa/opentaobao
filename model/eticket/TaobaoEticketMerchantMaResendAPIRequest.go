package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaResendAPIRequest
电子凭证重发回调接口 API请求
taobao.eticket.merchant.ma.resend

码商重发电子凭证回调接口 */
type TaobaoEticketMerchantMaResendAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 待重发的码列表
	_isvMaList []IsvMa
	// 业务id（订单号）
	_outerId string
	// 需要跟发码通知获取到的参数一致
	_token string
}

// NewTaobaoEticketMerchantMaResendRequest 初始化TaobaoEticketMerchantMaResendAPIRequest对象
func NewTaobaoEticketMerchantMaResendRequest() *TaobaoEticketMerchantMaResendAPIRequest {
	return &TaobaoEticketMerchantMaResendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaResendAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaResendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaResendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TaobaoEticketMerchantMaResendAPIRequest) GetBizType() int64 {
	return r._bizType
}

// Set is IsvMaList Setter
// 待重发的码列表
func (r *TaobaoEticketMerchantMaResendAPIRequest) SetIsvMaList(_isvMaList []IsvMa) error {
	r._isvMaList = _isvMaList
	r.Set("isv_ma_list", _isvMaList)
	return nil
}

// Get IsvMaList Getter
func (r TaobaoEticketMerchantMaResendAPIRequest) GetIsvMaList() []IsvMa {
	return r._isvMaList
}

// Set is OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaResendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoEticketMerchantMaResendAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaResendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r TaobaoEticketMerchantMaResendAPIRequest) GetToken() string {
	return r._token
}
