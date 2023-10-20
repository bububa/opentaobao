package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaReverseAPIRequest 电子凭证冲正接口 API请求
// taobao.eticket.merchant.ma.reverse
//
// 电子凭证平台冲正接口
type TaobaoEticketMerchantMaReverseAPIRequest struct {
	model.Params
	// 码值
	_code string
	// 业务id（订单号）
	_outerId string
	// 机具编号，如果核销时有则必传
	_posId string
	// 需要冲正的核销序列号
	_serialNum string
	// 需要跟发码通知获取到的参数一致
	_token string
	// 业务类型
	_bizType int64
	// 冲正份数，需要与核销份数一致
	_reverseNum int64
}

// NewTaobaoEticketMerchantMaReverseRequest 初始化TaobaoEticketMerchantMaReverseAPIRequest对象
func NewTaobaoEticketMerchantMaReverseRequest() *TaobaoEticketMerchantMaReverseAPIRequest {
	return &TaobaoEticketMerchantMaReverseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 码值
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetCode() string {
	return r._code
}

// SetOuterId is OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPosId is PosId Setter
// 机具编号，如果核销时有则必传
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetPosId(_posId string) error {
	r._posId = _posId
	r.Set("pos_id", _posId)
	return nil
}

// GetPosId PosId Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetPosId() string {
	return r._posId
}

// SetSerialNum is SerialNum Setter
// 需要冲正的核销序列号
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// GetSerialNum SerialNum Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetSerialNum() string {
	return r._serialNum
}

// SetToken is Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetToken() string {
	return r._token
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetReverseNum is ReverseNum Setter
// 冲正份数，需要与核销份数一致
func (r *TaobaoEticketMerchantMaReverseAPIRequest) SetReverseNum(_reverseNum int64) error {
	r._reverseNum = _reverseNum
	r.Set("reverse_num", _reverseNum)
	return nil
}

// GetReverseNum ReverseNum Getter
func (r TaobaoEticketMerchantMaReverseAPIRequest) GetReverseNum() int64 {
	return r._reverseNum
}
