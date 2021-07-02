package mozi

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziBucAccountPageallAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.buc.account.pageall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziBucAccountPageallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
