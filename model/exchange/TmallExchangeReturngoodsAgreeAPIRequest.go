package exchange

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeReturngoodsAgreeAPIRequest 卖家确认收货 API请求
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
type TmallExchangeReturngoodsAgreeAPIRequest struct {
	model.Params
	// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
	_fields []string
	// 换货单号ID
	_disputeId int64
}

// NewTmallExchangeReturngoodsAgreeRequest 初始化TmallExchangeReturngoodsAgreeAPIRequest对象
func NewTmallExchangeReturngoodsAgreeRequest() *TmallExchangeReturngoodsAgreeAPIRequest {
	return &TmallExchangeReturngoodsAgreeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallExchangeReturngoodsAgreeAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._disputeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.returngoods.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
func (r *TmallExchangeReturngoodsAgreeAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsAgreeAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

var poolTmallExchangeReturngoodsAgreeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallExchangeReturngoodsAgreeRequest()
	},
}

// GetTmallExchangeReturngoodsAgreeRequest 从 sync.Pool 获取 TmallExchangeReturngoodsAgreeAPIRequest
func GetTmallExchangeReturngoodsAgreeAPIRequest() *TmallExchangeReturngoodsAgreeAPIRequest {
	return poolTmallExchangeReturngoodsAgreeAPIRequest.Get().(*TmallExchangeReturngoodsAgreeAPIRequest)
}

// ReleaseTmallExchangeReturngoodsAgreeAPIRequest 将 TmallExchangeReturngoodsAgreeAPIRequest 放入 sync.Pool
func ReleaseTmallExchangeReturngoodsAgreeAPIRequest(v *TmallExchangeReturngoodsAgreeAPIRequest) {
	v.Reset()
	poolTmallExchangeReturngoodsAgreeAPIRequest.Put(v)
}
