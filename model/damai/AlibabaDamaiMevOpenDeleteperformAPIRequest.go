package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteperformAPIRequest 大麦换验平台-第三方对外开放-场次接口deletePerform API请求
// alibaba.damai.mev.open.deleteperform
//
// deletePerform
type AlibabaDamaiMevOpenDeleteperformAPIRequest struct {
	model.Params
	// 入参deletePerformParam
	_deletePerformParam *PerformIdOpenParam
}

// NewAlibabaDamaiMevOpenDeleteperformRequest 初始化AlibabaDamaiMevOpenDeleteperformAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteperformRequest() *AlibabaDamaiMevOpenDeleteperformAPIRequest {
	return &AlibabaDamaiMevOpenDeleteperformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenDeleteperformAPIRequest) Reset() {
	r._deletePerformParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteperform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeletePerformParam is DeletePerformParam Setter
// 入参deletePerformParam
func (r *AlibabaDamaiMevOpenDeleteperformAPIRequest) SetDeletePerformParam(_deletePerformParam *PerformIdOpenParam) error {
	r._deletePerformParam = _deletePerformParam
	r.Set("delete_perform_param", _deletePerformParam)
	return nil
}

// GetDeletePerformParam DeletePerformParam Getter
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetDeletePerformParam() *PerformIdOpenParam {
	return r._deletePerformParam
}

var poolAlibabaDamaiMevOpenDeleteperformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenDeleteperformRequest()
	},
}

// GetAlibabaDamaiMevOpenDeleteperformRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenDeleteperformAPIRequest
func GetAlibabaDamaiMevOpenDeleteperformAPIRequest() *AlibabaDamaiMevOpenDeleteperformAPIRequest {
	return poolAlibabaDamaiMevOpenDeleteperformAPIRequest.Get().(*AlibabaDamaiMevOpenDeleteperformAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenDeleteperformAPIRequest 将 AlibabaDamaiMevOpenDeleteperformAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeleteperformAPIRequest(v *AlibabaDamaiMevOpenDeleteperformAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeleteperformAPIRequest.Put(v)
}
