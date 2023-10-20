package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerQuerycapacitytaskAPIRequest 查询需求容量 API请求
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
type TmallServicecenterWorkerQuerycapacitytaskAPIRequest struct {
	model.Params
	// 查询对象
	_query *CapacityTaskQueryDto
}

// NewTmallServicecenterWorkerQuerycapacitytaskRequest 初始化TmallServicecenterWorkerQuerycapacitytaskAPIRequest对象
func NewTmallServicecenterWorkerQuerycapacitytaskRequest() *TmallServicecenterWorkerQuerycapacitytaskAPIRequest {
	return &TmallServicecenterWorkerQuerycapacitytaskAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkerQuerycapacitytaskAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQuerycapacitytaskAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.querycapacitytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQuerycapacitytaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkerQuerycapacitytaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询对象
func (r *TmallServicecenterWorkerQuerycapacitytaskAPIRequest) SetQuery(_query *CapacityTaskQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallServicecenterWorkerQuerycapacitytaskAPIRequest) GetQuery() *CapacityTaskQueryDto {
	return r._query
}

var poolTmallServicecenterWorkerQuerycapacitytaskAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkerQuerycapacitytaskRequest()
	},
}

// GetTmallServicecenterWorkerQuerycapacitytaskRequest 从 sync.Pool 获取 TmallServicecenterWorkerQuerycapacitytaskAPIRequest
func GetTmallServicecenterWorkerQuerycapacitytaskAPIRequest() *TmallServicecenterWorkerQuerycapacitytaskAPIRequest {
	return poolTmallServicecenterWorkerQuerycapacitytaskAPIRequest.Get().(*TmallServicecenterWorkerQuerycapacitytaskAPIRequest)
}

// ReleaseTmallServicecenterWorkerQuerycapacitytaskAPIRequest 将 TmallServicecenterWorkerQuerycapacitytaskAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkerQuerycapacitytaskAPIRequest(v *TmallServicecenterWorkerQuerycapacitytaskAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkerQuerycapacitytaskAPIRequest.Put(v)
}
