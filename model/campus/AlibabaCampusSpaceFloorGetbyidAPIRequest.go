package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceFloorGetbyidAPIRequest 根据id获取楼层 API请求
// alibaba.campus.space.floor.getbyid
//
// 根据id获取楼层
type AlibabaCampusSpaceFloorGetbyidAPIRequest struct {
	model.Params
	// 环境上下文
	_context *WorkBenchContext
	// 楼层id
	_id int64
}

// NewAlibabaCampusSpaceFloorGetbyidRequest 初始化AlibabaCampusSpaceFloorGetbyidAPIRequest对象
func NewAlibabaCampusSpaceFloorGetbyidRequest() *AlibabaCampusSpaceFloorGetbyidAPIRequest {
	return &AlibabaCampusSpaceFloorGetbyidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceFloorGetbyidAPIRequest) Reset() {
	r._context = nil
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.floor.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 环境上下文
func (r *AlibabaCampusSpaceFloorGetbyidAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetId is Id Setter
// 楼层id
func (r *AlibabaCampusSpaceFloorGetbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaCampusSpaceFloorGetbyidAPIRequest) GetId() int64 {
	return r._id
}

var poolAlibabaCampusSpaceFloorGetbyidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceFloorGetbyidRequest()
	},
}

// GetAlibabaCampusSpaceFloorGetbyidRequest 从 sync.Pool 获取 AlibabaCampusSpaceFloorGetbyidAPIRequest
func GetAlibabaCampusSpaceFloorGetbyidAPIRequest() *AlibabaCampusSpaceFloorGetbyidAPIRequest {
	return poolAlibabaCampusSpaceFloorGetbyidAPIRequest.Get().(*AlibabaCampusSpaceFloorGetbyidAPIRequest)
}

// ReleaseAlibabaCampusSpaceFloorGetbyidAPIRequest 将 AlibabaCampusSpaceFloorGetbyidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceFloorGetbyidAPIRequest(v *AlibabaCampusSpaceFloorGetbyidAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceFloorGetbyidAPIRequest.Put(v)
}
