package singletreasure

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询对象
func (r *TaobaoSingletreasureActivityQueryAPIRequest) SetQuery(_query *PageQueryDto) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoSingletreasureActivityQueryAPIRequest) GetQuery() *PageQueryDto {
	return r._query
}

var poolTaobaoSingletreasureActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityQueryRequest()
	},
}

// GetTaobaoSingletreasureActivityQueryRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityQueryAPIRequest
func GetTaobaoSingletreasureActivityQueryAPIRequest() *TaobaoSingletreasureActivityQueryAPIRequest {
	return poolTaobaoSingletreasureActivityQueryAPIRequest.Get().(*TaobaoSingletreasureActivityQueryAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityQueryAPIRequest 将 TaobaoSingletreasureActivityQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityQueryAPIRequest(v *TaobaoSingletreasureActivityQueryAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityQueryAPIRequest.Put(v)
}
