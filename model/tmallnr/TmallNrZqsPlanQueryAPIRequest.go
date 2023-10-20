package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrzqsplanqueryAPIRequest 周期送配送明细查询 API请求
// tmall.nr.zqs.plan.query
//
// 周期送配送明细查询
type TmallnrzqsplanqueryAPIRequest struct {
	model.Params
	// 交易子订单id
	_detailOrderId int64
}

// NewTmallnrzqsplanqueryRequest 初始化TmallnrzqsplanqueryAPIRequest对象
func NewTmallnrzqsplanqueryRequest() *TmallnrzqsplanqueryAPIRequest {
	return &TmallnrzqsplanqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrzqsplanqueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.zqs.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrzqsplanqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrzqsplanqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailOrderId is DetailOrderId Setter
// 交易子订单id
func (r *TmallnrzqsplanqueryAPIRequest) SetDetailOrderId(_detailOrderId int64) error {
	r._detailOrderId = _detailOrderId
	r.Set("detail_order_id", _detailOrderId)
	return nil
}

// GetDetailOrderId DetailOrderId Getter
func (r TmallnrzqsplanqueryAPIRequest) GetDetailOrderId() int64 {
	return r._detailOrderId
}
