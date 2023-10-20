package eticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaSendAPIRequest 码商发码成功回调接口 API请求
// taobao.eticket.merchant.ma.send
//
// 码商发码成功回调接口
type TaobaoEticketMerchantMaSendAPIRequest struct {
	model.Params
	// 需要发送的码列表
	_isvMaList []IsvMa
	// 业务id（订单号）
	_outerId string
	// 需要跟发码通知获取到的参数一致
	_token string
	// 业务类型
	_bizType int64
}

// NewTaobaoEticketMerchantMaSendRequest 初始化TaobaoEticketMerchantMaSendAPIRequest对象
func NewTaobaoEticketMerchantMaSendRequest() *TaobaoEticketMerchantMaSendAPIRequest {
	return &TaobaoEticketMerchantMaSendAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoEticketMerchantMaSendAPIRequest) Reset() {
	r._isvMaList = r._isvMaList[:0]
	r._outerId = ""
	r._token = ""
	r._bizType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaSendAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoEticketMerchantMaSendAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoEticketMerchantMaSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoEticketMerchantMaSendRequest()
	},
}

// GetTaobaoEticketMerchantMaSendRequest 从 sync.Pool 获取 TaobaoEticketMerchantMaSendAPIRequest
func GetTaobaoEticketMerchantMaSendAPIRequest() *TaobaoEticketMerchantMaSendAPIRequest {
	return poolTaobaoEticketMerchantMaSendAPIRequest.Get().(*TaobaoEticketMerchantMaSendAPIRequest)
}

// ReleaseTaobaoEticketMerchantMaSendAPIRequest 将 TaobaoEticketMerchantMaSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoEticketMerchantMaSendAPIRequest(v *TaobaoEticketMerchantMaSendAPIRequest) {
	v.Reset()
	poolTaobaoEticketMerchantMaSendAPIRequest.Put(v)
}
