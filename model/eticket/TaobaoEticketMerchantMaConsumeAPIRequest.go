package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchantmaconsumeAPIRequest 电子凭证核销接口 API请求
// taobao.eticket.merchant.ma.consume
//
// 电子凭证核销接口
type TaobaoeticketmerchantmaconsumeAPIRequest struct {
	model.Params
	// 核销后换码的码列表
	_isvMaList []IsvMa
	// 需要被核销的码
	_code string
	// 业务id（订单号）
	_outerId string
	// 机具编号
	_posId string
	// 核销序列号，需要保证唯一
	_serialNum string
	// 需要跟发码通知获取到的参数一致
	_token string
	// 业务类型
	_bizType int64
	// 核销份数
	_consumeNum int64
}

// NewTaobaoeticketmerchantmaconsumeRequest 初始化TaobaoeticketmerchantmaconsumeAPIRequest对象
func NewTaobaoeticketmerchantmaconsumeRequest() *TaobaoeticketmerchantmaconsumeAPIRequest {
	return &TaobaoeticketmerchantmaconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvMaList is IsvMaList Setter
// 核销后换码的码列表
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetIsvMaList(_isvMaList []IsvMa) error {
	r._isvMaList = _isvMaList
	r.Set("isv_ma_list", _isvMaList)
	return nil
}

// GetIsvMaList IsvMaList Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetIsvMaList() []IsvMa {
	return r._isvMaList
}

// SetCode is Code Setter
// 需要被核销的码
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetCode() string {
	return r._code
}

// SetOuterId is OuterId Setter
// 业务id（订单号）
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPosId is PosId Setter
// 机具编号
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetPosId(_posId string) error {
	r._posId = _posId
	r.Set("pos_id", _posId)
	return nil
}

// GetPosId PosId Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetPosId() string {
	return r._posId
}

// SetSerialNum is SerialNum Setter
// 核销序列号，需要保证唯一
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// GetSerialNum SerialNum Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetSerialNum() string {
	return r._serialNum
}

// SetToken is Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetToken() string {
	return r._token
}

// SetBizType is BizType Setter
// 业务类型
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetConsumeNum is ConsumeNum Setter
// 核销份数
func (r *TaobaoeticketmerchantmaconsumeAPIRequest) SetConsumeNum(_consumeNum int64) error {
	r._consumeNum = _consumeNum
	r.Set("consume_num", _consumeNum)
	return nil
}

// GetConsumeNum ConsumeNum Getter
func (r TaobaoeticketmerchantmaconsumeAPIRequest) GetConsumeNum() int64 {
	return r._consumeNum
}
