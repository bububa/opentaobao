package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchantmafailsendAPIRequest 码商发码失败回调接口 API请求
// taobao.eticket.merchant.ma.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
type TaobaoeticketmerchantmafailsendAPIRequest struct {
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

// NewTaobaoeticketmerchantmafailsendRequest 初始化TaobaoeticketmerchantmafailsendAPIRequest对象
func NewTaobaoeticketmerchantmafailsendRequest() *TaobaoeticketmerchantmafailsendAPIRequest {
	return &TaobaoeticketmerchantmafailsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.failsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 业务id（订单号）
func (r *TaobaoeticketmerchantmafailsendAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetSubErrCode is SubErrCode Setter
// 错误原因码
func (r *TaobaoeticketmerchantmafailsendAPIRequest) SetSubErrCode(_subErrCode string) error {
	r._subErrCode = _subErrCode
	r.Set("sub_err_code", _subErrCode)
	return nil
}

// GetSubErrCode SubErrCode Getter
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetSubErrCode() string {
	return r._subErrCode
}

// SetSubErrMsg is SubErrMsg Setter
// 错误码描述
func (r *TaobaoeticketmerchantmafailsendAPIRequest) SetSubErrMsg(_subErrMsg string) error {
	r._subErrMsg = _subErrMsg
	r.Set("sub_err_msg", _subErrMsg)
	return nil
}

// GetSubErrMsg SubErrMsg Getter
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetSubErrMsg() string {
	return r._subErrMsg
}

// SetToken is Token Setter
// 需要与发码通知获取的值一致
func (r *TaobaoeticketmerchantmafailsendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetToken() string {
	return r._token
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoeticketmerchantmafailsendAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoeticketmerchantmafailsendAPIRequest) GetBizType() int64 {
	return r._bizType
}
