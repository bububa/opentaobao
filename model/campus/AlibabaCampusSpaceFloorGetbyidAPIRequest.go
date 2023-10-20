package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacefloorgetbyidAPIRequest 根据id获取楼层 API请求
// alibaba.campus.space.floor.getbyid
//
// 根据id获取楼层
type AlibabacampusspacefloorgetbyidAPIRequest struct {
	model.Params
	// 环境上下文
	_context *WorkBenchContext
	// 楼层id
	_id int64
}

// NewAlibabacampusspacefloorgetbyidRequest 初始化AlibabacampusspacefloorgetbyidAPIRequest对象
func NewAlibabacampusspacefloorgetbyidRequest() *AlibabacampusspacefloorgetbyidAPIRequest {
	return &AlibabacampusspacefloorgetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacefloorgetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.floor.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacefloorgetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacefloorgetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 环境上下文
func (r *AlibabacampusspacefloorgetbyidAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusspacefloorgetbyidAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetId is Id Setter
// 楼层id
func (r *AlibabacampusspacefloorgetbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabacampusspacefloorgetbyidAPIRequest) GetId() int64 {
	return r._id
}
