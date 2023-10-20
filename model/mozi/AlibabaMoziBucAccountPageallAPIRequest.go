package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozibucaccountpageallAPIRequest 查询租户内内所有账号 API请求
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
type AlibabamozibucaccountpageallAPIRequest struct {
	model.Params
	// 查询租户内所有人员和账号
	_pageAll *PageAllAccountsRequest
}

// NewAlibabamozibucaccountpageallRequest 初始化AlibabamozibucaccountpageallAPIRequest对象
func NewAlibabamozibucaccountpageallRequest() *AlibabamozibucaccountpageallAPIRequest {
	return &AlibabamozibucaccountpageallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozibucaccountpageallAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.buc.account.pageall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozibucaccountpageallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozibucaccountpageallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageAll is PageAll Setter
// 查询租户内所有人员和账号
func (r *AlibabamozibucaccountpageallAPIRequest) SetPageAll(_pageAll *PageAllAccountsRequest) error {
	r._pageAll = _pageAll
	r.Set("page_all", _pageAll)
	return nil
}

// GetPageAll PageAll Getter
func (r AlibabamozibucaccountpageallAPIRequest) GetPageAll() *PageAllAccountsRequest {
	return r._pageAll
}
