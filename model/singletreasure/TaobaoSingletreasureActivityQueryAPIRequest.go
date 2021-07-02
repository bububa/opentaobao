package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityQueryAPIRequest 查询活动列表接口 API请求
// taobao.singletreasure.activity.query
//
// 查询活动列表接口
type TaobaoSingletreasureActivityQueryAPIRequest struct {
	model.Params
	// 查询对象
	_query *PageQueryDto
}

// NewTaobaoSingletreasureActivityQueryRequest 初始化TaobaoSingletreasureActivityQueryAPIRequest对象
func NewTaobaoSingletreasureActivityQueryRequest() *TaobaoSingletreasureActivityQueryAPIRequest {
	return &TaobaoSingletreasureActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 查询对象
func (r *TaobaoSingletreasureActivityQueryAPIRequest) SetQuery(_query *PageQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetQuery() *PageQueryDto {
	return r._query
}
