package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityqueryAPIRequest 查询活动列表接口 API请求
// taobao.singletreasure.activity.query
//
// 查询活动列表接口
type TaobaosingletreasureactivityqueryAPIRequest struct {
	model.Params
	// 查询对象
	_query *PageQueryDto
}

// NewTaobaosingletreasureactivityqueryRequest 初始化TaobaosingletreasureactivityqueryAPIRequest对象
func NewTaobaosingletreasureactivityqueryRequest() *TaobaosingletreasureactivityqueryAPIRequest {
	return &TaobaosingletreasureactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivityqueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询对象
func (r *TaobaosingletreasureactivityqueryAPIRequest) SetQuery(_query *PageQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaosingletreasureactivityqueryAPIRequest) GetQuery() *PageQueryDto {
	return r._query
}
