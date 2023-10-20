package aliexpresssumaitong

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeOrderOpenQueryAPIRequest Aliexpress开放平台订单查询 API请求
// aliexpress.trade.order.open.query
//
// Aliexpress开放平台订单信息查询
type AliexpressTradeOrderOpenQueryAPIRequest struct {
	model.Params
	// 订单号
	_orderIds []int64
	// 外部订单号
	_outIds []string
	// 小程序appId
	_openAppKey string
	// 业务编码
	_bizCode string
	// 买家用户id
	_buyerId int64
}

// NewAliexpressTradeOrderOpenQueryRequest 初始化AliexpressTradeOrderOpenQueryAPIRequest对象
func NewAliexpressTradeOrderOpenQueryRequest() *AliexpressTradeOrderOpenQueryAPIRequest {
	return &AliexpressTradeOrderOpenQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressTradeOrderOpenQueryAPIRequest) Reset() {
	r._orderIds = r._orderIds[:0]
	r._outIds = r._outIds[:0]
	r._openAppKey = ""
	r._bizCode = ""
	r._buyerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.order.open.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIds is OrderIds Setter
// 订单号
func (r *AliexpressTradeOrderOpenQueryAPIRequest) SetOrderIds(_orderIds []int64) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// GetOrderIds OrderIds Getter
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetOrderIds() []int64 {
	return r._orderIds
}

// SetOutIds is OutIds Setter
// 外部订单号
func (r *AliexpressTradeOrderOpenQueryAPIRequest) SetOutIds(_outIds []string) error {
	r._outIds = _outIds
	r.Set("out_ids", _outIds)
	return nil
}

// GetOutIds OutIds Getter
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetOutIds() []string {
	return r._outIds
}

// SetOpenAppKey is OpenAppKey Setter
// 小程序appId
func (r *AliexpressTradeOrderOpenQueryAPIRequest) SetOpenAppKey(_openAppKey string) error {
	r._openAppKey = _openAppKey
	r.Set("open_app_key", _openAppKey)
	return nil
}

// GetOpenAppKey OpenAppKey Getter
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetOpenAppKey() string {
	return r._openAppKey
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *AliexpressTradeOrderOpenQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetBuyerId is BuyerId Setter
// 买家用户id
func (r *AliexpressTradeOrderOpenQueryAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r AliexpressTradeOrderOpenQueryAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}

var poolAliexpressTradeOrderOpenQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressTradeOrderOpenQueryRequest()
	},
}

// GetAliexpressTradeOrderOpenQueryRequest 从 sync.Pool 获取 AliexpressTradeOrderOpenQueryAPIRequest
func GetAliexpressTradeOrderOpenQueryAPIRequest() *AliexpressTradeOrderOpenQueryAPIRequest {
	return poolAliexpressTradeOrderOpenQueryAPIRequest.Get().(*AliexpressTradeOrderOpenQueryAPIRequest)
}

// ReleaseAliexpressTradeOrderOpenQueryAPIRequest 将 AliexpressTradeOrderOpenQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressTradeOrderOpenQueryAPIRequest(v *AliexpressTradeOrderOpenQueryAPIRequest) {
	v.Reset()
	poolAliexpressTradeOrderOpenQueryAPIRequest.Put(v)
}
