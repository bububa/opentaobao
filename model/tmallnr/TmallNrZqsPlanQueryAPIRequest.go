package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrZqsPlanQueryAPIRequest 周期送配送明细查询 API请求
// tmall.nr.zqs.plan.query
//
// 周期送配送明细查询
type TmallNrZqsPlanQueryAPIRequest struct {
	model.Params
	// 交易子订单id
	_detailOrderId int64
}

// NewTmallNrZqsPlanQueryRequest 初始化TmallNrZqsPlanQueryAPIRequest对象
func NewTmallNrZqsPlanQueryRequest() *TmallNrZqsPlanQueryAPIRequest {
	return &TmallNrZqsPlanQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrZqsPlanQueryAPIRequest) Reset() {
	r._detailOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrZqsPlanQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nr.zqs.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrZqsPlanQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrZqsPlanQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailOrderId is DetailOrderId Setter
// 交易子订单id
func (r *TmallNrZqsPlanQueryAPIRequest) SetDetailOrderId(_detailOrderId int64) error {
	r._detailOrderId = _detailOrderId
	r.Set("detail_order_id", _detailOrderId)
	return nil
}

// GetDetailOrderId DetailOrderId Getter
func (r TmallNrZqsPlanQueryAPIRequest) GetDetailOrderId() int64 {
	return r._detailOrderId
}

var poolTmallNrZqsPlanQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrZqsPlanQueryRequest()
	},
}

// GetTmallNrZqsPlanQueryRequest 从 sync.Pool 获取 TmallNrZqsPlanQueryAPIRequest
func GetTmallNrZqsPlanQueryAPIRequest() *TmallNrZqsPlanQueryAPIRequest {
	return poolTmallNrZqsPlanQueryAPIRequest.Get().(*TmallNrZqsPlanQueryAPIRequest)
}

// ReleaseTmallNrZqsPlanQueryAPIRequest 将 TmallNrZqsPlanQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrZqsPlanQueryAPIRequest(v *TmallNrZqsPlanQueryAPIRequest) {
	v.Reset()
	poolTmallNrZqsPlanQueryAPIRequest.Put(v)
}
