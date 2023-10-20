package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkerquerycapacitytaskAPIRequest 查询需求容量 API请求
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
type TmallservicecenterworkerquerycapacitytaskAPIRequest struct {
	model.Params
	// 查询对象
	_query *CapacityTaskQueryDto
}

// NewTmallservicecenterworkerquerycapacitytaskRequest 初始化TmallservicecenterworkerquerycapacitytaskAPIRequest对象
func NewTmallservicecenterworkerquerycapacitytaskRequest() *TmallservicecenterworkerquerycapacitytaskAPIRequest {
	return &TmallservicecenterworkerquerycapacitytaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkerquerycapacitytaskAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.querycapacitytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkerquerycapacitytaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkerquerycapacitytaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询对象
func (r *TmallservicecenterworkerquerycapacitytaskAPIRequest) SetQuery(_query *CapacityTaskQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallservicecenterworkerquerycapacitytaskAPIRequest) GetQuery() *CapacityTaskQueryDto {
	return r._query
}
