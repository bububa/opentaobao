package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaFailsendAPIRequest
码商发码失败回调接口 API请求
taobao.eticket.merchant.ma.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证 */
type TaobaoEticketMerchantMaFailsendAPIRequest struct {
	model.Params
	// 业务id（订单号）
	_outerId string
	// 错误原因码
	_subErrCode string
	// 错误码描述
	_subErrMsg string
	// 需要与发码通知获取的值一致
	_token string
	// 业务类型
	_bizType int64
}

// NewTaobaoEticketMerchantMaFailsendRequest 初始化TaobaoEticketMerchantMaFailsendAPIRequest对象
func NewTaobaoEticketMerchantMaFailsendRequest() *TaobaoEticketMerchantMaFailsendAPIRequest {
	return &TaobaoEticketMerchantMaFailsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.failsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaFailsendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is SubErrCode Setter
// 错误原因码
func (r *TaobaoEticketMerchantMaFailsendAPIRequest) SetSubErrCode(_subErrCode string) error {
	r._subErrCode = _subErrCode
	r.Set("sub_err_code", _subErrCode)
	return nil
}

// Get SubErrCode Getter
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetSubErrCode() string {
	return r._subErrCode
}

// Set is SubErrMsg Setter
// 错误码描述
func (r *TaobaoEticketMerchantMaFailsendAPIRequest) SetSubErrMsg(_subErrMsg string) error {
	r._subErrMsg = _subErrMsg
	r.Set("sub_err_msg", _subErrMsg)
	return nil
}

// Get SubErrMsg Getter
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetSubErrMsg() string {
	return r._subErrMsg
}

// Set is Token Setter
// 需要与发码通知获取的值一致
func (r *TaobaoEticketMerchantMaFailsendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetToken() string {
	return r._token
}

// Set is BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaFailsendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TaobaoEticketMerchantMaFailsendAPIRequest) GetBizType() int64 {
	return r._bizType
}
