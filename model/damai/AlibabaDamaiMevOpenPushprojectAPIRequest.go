package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushprojectAPIRequest 大麦换验平台-第三方对外开放-项目接口pushProject API请求
// alibaba.damai.mev.open.pushproject
//
// pushProject
type AlibabaDamaiMevOpenPushprojectAPIRequest struct {
	model.Params
	// 入参pushProjectParam
	_pushProjectParam *ThirdProjectPushOpenParam
}

// NewAlibabaDamaiMevOpenPushprojectRequest 初始化AlibabaDamaiMevOpenPushprojectAPIRequest对象
func NewAlibabaDamaiMevOpenPushprojectRequest() *AlibabaDamaiMevOpenPushprojectAPIRequest {
	return &AlibabaDamaiMevOpenPushprojectAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushprojectAPIRequest) Reset() {
	r._pushProjectParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.pushproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushProjectParam is PushProjectParam Setter
// 入参pushProjectParam
func (r *AlibabaDamaiMevOpenPushprojectAPIRequest) SetPushProjectParam(_pushProjectParam *ThirdProjectPushOpenParam) error {
	r._pushProjectParam = _pushProjectParam
	r.Set("push_project_param", _pushProjectParam)
	return nil
}

// GetPushProjectParam PushProjectParam Getter
func (r AlibabaDamaiMevOpenPushprojectAPIRequest) GetPushProjectParam() *ThirdProjectPushOpenParam {
	return r._pushProjectParam
}

var poolAlibabaDamaiMevOpenPushprojectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushprojectRequest()
	},
}

// GetAlibabaDamaiMevOpenPushprojectRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushprojectAPIRequest
func GetAlibabaDamaiMevOpenPushprojectAPIRequest() *AlibabaDamaiMevOpenPushprojectAPIRequest {
	return poolAlibabaDamaiMevOpenPushprojectAPIRequest.Get().(*AlibabaDamaiMevOpenPushprojectAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushprojectAPIRequest 将 AlibabaDamaiMevOpenPushprojectAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushprojectAPIRequest(v *AlibabaDamaiMevOpenPushprojectAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushprojectAPIRequest.Put(v)
}
