package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeQueryAPIRequest 新商场当面付交易查询 API请求
// alibaba.mos.onsite.trade.query
//
// 本接口提供新商场当面付订单的查询的功能，商户可以通过本接口主动查询订单状态，完成下一步的业务逻辑。
// 商户系统应在两种场景下调用此接口: 商户POS系统应该在调用[条码支付请求接口]并成功返回后，调用此接口查询订单的处理状态。
type AlibabaMosOnsiteTradeQueryAPIRequest struct {
	model.Params
	// 喵街交易流水号。与商户支付流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号  > 商户支付流水号。
	_tradeNo string
	// 原支付请求的商户支付流水号。与喵街交易流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号 >  商户支付流水号。
	_outTradeNo string
	// 商户ID类型，取值为miaojie或out
	_storeIdType string
	// 门店号或喵街商户ID
	_storeId string
}

// NewAlibabaMosOnsiteTradeQueryRequest 初始化AlibabaMosOnsiteTradeQueryAPIRequest对象
func NewAlibabaMosOnsiteTradeQueryRequest() *AlibabaMosOnsiteTradeQueryAPIRequest {
	return &AlibabaMosOnsiteTradeQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosOnsiteTradeQueryAPIRequest) Reset() {
	r._tradeNo = ""
	r._outTradeNo = ""
	r._storeIdType = ""
	r._storeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 喵街交易流水号。与商户支付流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号  &gt; 商户支付流水号。
func (r *AlibabaMosOnsiteTradeQueryAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutTradeNo is OutTradeNo Setter
// 原支付请求的商户支付流水号。与喵街交易流水号两者至少要填写一个。如果均有，优先级为喵街交易流水号 &gt;  商户支付流水号。
func (r *AlibabaMosOnsiteTradeQueryAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreIdType is StoreIdType Setter
// 商户ID类型，取值为miaojie或out
func (r *AlibabaMosOnsiteTradeQueryAPIRequest) SetStoreIdType(_storeIdType string) error {
	r._storeIdType = _storeIdType
	r.Set("store_id_type", _storeIdType)
	return nil
}

// GetStoreIdType StoreIdType Getter
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetStoreIdType() string {
	return r._storeIdType
}

// SetStoreId is StoreId Setter
// 门店号或喵街商户ID
func (r *AlibabaMosOnsiteTradeQueryAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaMosOnsiteTradeQueryAPIRequest) GetStoreId() string {
	return r._storeId
}

var poolAlibabaMosOnsiteTradeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosOnsiteTradeQueryRequest()
	},
}

// GetAlibabaMosOnsiteTradeQueryRequest 从 sync.Pool 获取 AlibabaMosOnsiteTradeQueryAPIRequest
func GetAlibabaMosOnsiteTradeQueryAPIRequest() *AlibabaMosOnsiteTradeQueryAPIRequest {
	return poolAlibabaMosOnsiteTradeQueryAPIRequest.Get().(*AlibabaMosOnsiteTradeQueryAPIRequest)
}

// ReleaseAlibabaMosOnsiteTradeQueryAPIRequest 将 AlibabaMosOnsiteTradeQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosOnsiteTradeQueryAPIRequest(v *AlibabaMosOnsiteTradeQueryAPIRequest) {
	v.Reset()
	poolAlibabaMosOnsiteTradeQueryAPIRequest.Put(v)
}
