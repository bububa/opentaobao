package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaSendAPIRequest 码商发码成功回调接口 API请求
// taobao.eticket.merchant.ma.send
//
// 码商发码成功回调接口
type TaobaoEticketMerchantMaSendAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 需要发送的码列表
	_isvMaList []IsvMa
	// 业务id（订单号）
	_outerId string
	// 需要跟发码通知获取到的参数一致
	_token string
}

// NewTaobaoEticketMerchantMaSendRequest 初始化TaobaoEticketMerchantMaSendAPIRequest对象
func NewTaobaoEticketMerchantMaSendRequest() *TaobaoEticketMerchantMaSendAPIRequest {
	return &TaobaoEticketMerchantMaSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaSendAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaSendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoEticketMerchantMaSendAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetIsvMaList is IsvMaList Setter
// 需要发送的码列表
func (r *TaobaoEticketMerchantMaSendAPIRequest) SetIsvMaList(_isvMaList []IsvMa) error {
	r._isvMaList = _isvMaList
	r.Set("isv_ma_list", _isvMaList)
	return nil
}

// GetIsvMaList IsvMaList Getter
func (r TaobaoEticketMerchantMaSendAPIRequest) GetIsvMaList() []IsvMa {
	return r._isvMaList
}

// SetOuterId is OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaSendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoEticketMerchantMaSendAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetToken is Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaSendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoEticketMerchantMaSendAPIRequest) GetToken() string {
	return r._token
}
