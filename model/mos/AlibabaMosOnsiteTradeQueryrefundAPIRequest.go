package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosonsitetradequeryrefundAPIRequest 退款查询 API请求
// alibaba.mos.onsite.trade.queryrefund
//
// 商户可使用该接口查询退款请求是否执行成功。
type AlibabamosonsitetradequeryrefundAPIRequest struct {
	model.Params
	// 订单号。可能为外部订单号，也可能为喵街订单号
	_orderNo string
	// 退款外部流水号
	_outRequestNo string
}

// NewAlibabamosonsitetradequeryrefundRequest 初始化AlibabamosonsitetradequeryrefundAPIRequest对象
func NewAlibabamosonsitetradequeryrefundRequest() *AlibabamosonsitetradequeryrefundAPIRequest {
	return &AlibabamosonsitetradequeryrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosonsitetradequeryrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.queryrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosonsitetradequeryrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosonsitetradequeryrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单号。可能为外部订单号，也可能为喵街订单号
func (r *AlibabamosonsitetradequeryrefundAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabamosonsitetradequeryrefundAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetOutRequestNo is OutRequestNo Setter
// 退款外部流水号
func (r *AlibabamosonsitetradequeryrefundAPIRequest) SetOutRequestNo(_outRequestNo string) error {
	r._outRequestNo = _outRequestNo
	r.Set("out_request_no", _outRequestNo)
	return nil
}

// GetOutRequestNo OutRequestNo Getter
func (r AlibabamosonsitetradequeryrefundAPIRequest) GetOutRequestNo() string {
	return r._outRequestNo
}
