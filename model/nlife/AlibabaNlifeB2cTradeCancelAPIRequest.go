package nlife

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeCancelAPIRequest 零售+平台取消订单 API请求
// alibaba.nlife.b2c.trade.cancel
//
// 零售+平台取消订单接口
type AlibabaNlifeB2cTradeCancelAPIRequest struct {
	model.Params
	// 零售+平台订单号，和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店号
	_storeId string
}

// NewAlibabaNlifeB2cTradeCancelRequest 初始化AlibabaNlifeB2cTradeCancelAPIRequest对象
func NewAlibabaNlifeB2cTradeCancelRequest() *AlibabaNlifeB2cTradeCancelAPIRequest {
	return &AlibabaNlifeB2cTradeCancelAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) Reset() {
	r._tradeNo = ""
	r._outTradeNo = ""
	r._storeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreId is StoreId Setter
// 零售+门店号
func (r *AlibabaNlifeB2cTradeCancelAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaNlifeB2cTradeCancelAPIRequest) GetStoreId() string {
	return r._storeId
}

var poolAlibabaNlifeB2cTradeCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNlifeB2cTradeCancelRequest()
	},
}

// GetAlibabaNlifeB2cTradeCancelRequest 从 sync.Pool 获取 AlibabaNlifeB2cTradeCancelAPIRequest
func GetAlibabaNlifeB2cTradeCancelAPIRequest() *AlibabaNlifeB2cTradeCancelAPIRequest {
	return poolAlibabaNlifeB2cTradeCancelAPIRequest.Get().(*AlibabaNlifeB2cTradeCancelAPIRequest)
}

// ReleaseAlibabaNlifeB2cTradeCancelAPIRequest 将 AlibabaNlifeB2cTradeCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaNlifeB2cTradeCancelAPIRequest(v *AlibabaNlifeB2cTradeCancelAPIRequest) {
	v.Reset()
	poolAlibabaNlifeB2cTradeCancelAPIRequest.Put(v)
}
