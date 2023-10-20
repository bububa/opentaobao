package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziBucAccountPageallAPIRequest 查询租户内内所有账号 API请求
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
type AlibabaMoziBucAccountPageallAPIRequest struct {
	model.Params
	// 查询租户内所有人员和账号
	_pageAll *PageAllAccountsRequest
}

// NewAlibabaMoziBucAccountPageallRequest 初始化AlibabaMoziBucAccountPageallAPIRequest对象
func NewAlibabaMoziBucAccountPageallRequest() *AlibabaMoziBucAccountPageallAPIRequest {
	return &AlibabaMoziBucAccountPageallAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziBucAccountPageallAPIRequest) Reset() {
	r._pageAll = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziBucAccountPageallAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.buc.account.pageall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziBucAccountPageallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziBucAccountPageallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageAll is PageAll Setter
// 查询租户内所有人员和账号
func (r *AlibabaMoziBucAccountPageallAPIRequest) SetPageAll(_pageAll *PageAllAccountsRequest) error {
	r._pageAll = _pageAll
	r.Set("page_all", _pageAll)
	return nil
}

// GetPageAll PageAll Getter
func (r AlibabaMoziBucAccountPageallAPIRequest) GetPageAll() *PageAllAccountsRequest {
	return r._pageAll
}

var poolAlibabaMoziBucAccountPageallAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziBucAccountPageallRequest()
	},
}

// GetAlibabaMoziBucAccountPageallRequest 从 sync.Pool 获取 AlibabaMoziBucAccountPageallAPIRequest
func GetAlibabaMoziBucAccountPageallAPIRequest() *AlibabaMoziBucAccountPageallAPIRequest {
	return poolAlibabaMoziBucAccountPageallAPIRequest.Get().(*AlibabaMoziBucAccountPageallAPIRequest)
}

// ReleaseAlibabaMoziBucAccountPageallAPIRequest 将 AlibabaMoziBucAccountPageallAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziBucAccountPageallAPIRequest(v *AlibabaMoziBucAccountPageallAPIRequest) {
	v.Reset()
	poolAlibabaMoziBucAccountPageallAPIRequest.Put(v)
}
