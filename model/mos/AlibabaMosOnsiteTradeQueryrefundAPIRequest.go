package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeQueryrefundAPIRequest 退款查询 API请求
// alibaba.mos.onsite.trade.queryrefund
//
// 商户可使用该接口查询退款请求是否执行成功。
type AlibabaMosOnsiteTradeQueryrefundAPIRequest struct {
	model.Params
	// 订单号。可能为外部订单号，也可能为喵街订单号
	_orderNo string
	// 退款外部流水号
	_outRequestNo string
}

// NewAlibabaMosOnsiteTradeQueryrefundRequest 初始化AlibabaMosOnsiteTradeQueryrefundAPIRequest对象
func NewAlibabaMosOnsiteTradeQueryrefundRequest() *AlibabaMosOnsiteTradeQueryrefundAPIRequest {
	return &AlibabaMosOnsiteTradeQueryrefundAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosOnsiteTradeQueryrefundAPIRequest) Reset() {
	r._orderNo = ""
	r._outRequestNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeQueryrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.queryrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeQueryrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosOnsiteTradeQueryrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单号。可能为外部订单号，也可能为喵街订单号
func (r *AlibabaMosOnsiteTradeQueryrefundAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabaMosOnsiteTradeQueryrefundAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetOutRequestNo is OutRequestNo Setter
// 退款外部流水号
func (r *AlibabaMosOnsiteTradeQueryrefundAPIRequest) SetOutRequestNo(_outRequestNo string) error {
	r._outRequestNo = _outRequestNo
	r.Set("out_request_no", _outRequestNo)
	return nil
}

// GetOutRequestNo OutRequestNo Getter
func (r AlibabaMosOnsiteTradeQueryrefundAPIRequest) GetOutRequestNo() string {
	return r._outRequestNo
}

var poolAlibabaMosOnsiteTradeQueryrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosOnsiteTradeQueryrefundRequest()
	},
}

// GetAlibabaMosOnsiteTradeQueryrefundRequest 从 sync.Pool 获取 AlibabaMosOnsiteTradeQueryrefundAPIRequest
func GetAlibabaMosOnsiteTradeQueryrefundAPIRequest() *AlibabaMosOnsiteTradeQueryrefundAPIRequest {
	return poolAlibabaMosOnsiteTradeQueryrefundAPIRequest.Get().(*AlibabaMosOnsiteTradeQueryrefundAPIRequest)
}

// ReleaseAlibabaMosOnsiteTradeQueryrefundAPIRequest 将 AlibabaMosOnsiteTradeQueryrefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosOnsiteTradeQueryrefundAPIRequest(v *AlibabaMosOnsiteTradeQueryrefundAPIRequest) {
	v.Reset()
	poolAlibabaMosOnsiteTradeQueryrefundAPIRequest.Put(v)
}
