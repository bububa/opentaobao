package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 API请求
// alibaba.campus.space.unit.getlistwithattrbygroupid
//
// 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 分组id
	_groupId int64
}

// NewAlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest 初始化AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest对象
func NewAlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest() *AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest {
	return &AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) Reset() {
	r._context = nil
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlistwithattrbygroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetGroupId is GroupId Setter
// 分组id
func (r *AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) GetGroupId() int64 {
	return r._groupId
}

var poolAlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest()
	},
}

// GetAlibabaCampusSpaceUnitGetlistwithattrbygroupidRequest 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest
func GetAlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest() *AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest {
	return poolAlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest.Get().(*AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest)
}

// ReleaseAlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest 将 AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest(v *AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest.Put(v)
}
