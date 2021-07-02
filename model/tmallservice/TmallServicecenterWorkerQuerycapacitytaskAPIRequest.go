package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQuerycapacitytaskAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.querycapacitytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQuerycapacitytaskAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
