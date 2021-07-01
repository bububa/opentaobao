package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrZqsPlanQueryAPIRequest
周期送配送明细查询 API请求
tmall.nr.zqs.plan.query

周期送配送明细查询 */
type TmallNrZqsPlanQueryAPIRequest struct {
	model.Params
	// 交易子订单id
	_detailOrderId int64
}

// NewTmallNrZqsPlanQueryRequest 初始化TmallNrZqsPlanQueryAPIRequest对象
func NewTmallNrZqsPlanQueryRequest() *TmallNrZqsPlanQueryAPIRequest {
	return &TmallNrZqsPlanQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrZqsPlanQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.zqs.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrZqsPlanQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DetailOrderId Setter
// 交易子订单id
func (r *TmallNrZqsPlanQueryAPIRequest) SetDetailOrderId(_detailOrderId int64) error {
	r._detailOrderId = _detailOrderId
	r.Set("detail_order_id", _detailOrderId)
	return nil
}

// Get DetailOrderId Getter
func (r TmallNrZqsPlanQueryAPIRequest) GetDetailOrderId() int64 {
	return r._detailOrderId
}
