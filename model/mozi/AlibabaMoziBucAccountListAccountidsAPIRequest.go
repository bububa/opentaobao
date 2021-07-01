package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziBucAccountListAccountidsAPIRequest
根据一批账号ID查询账号列表 API请求
alibaba.mozi.buc.account.list.accountids

根据一批账号ID查询账号列表 */
type AlibabaMoziBucAccountListAccountidsAPIRequest struct {
	model.Params
	// 请求参数
	_listAccountIds *ListAccountsByAccountIdsRequest
}

// NewAlibabaMoziBucAccountListAccountidsRequest 初始化AlibabaMoziBucAccountListAccountidsAPIRequest对象
func NewAlibabaMoziBucAccountListAccountidsRequest() *AlibabaMoziBucAccountListAccountidsAPIRequest {
	return &AlibabaMoziBucAccountListAccountidsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.buc.account.list.accountids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ListAccountIds Setter
// 请求参数
func (r *AlibabaMoziBucAccountListAccountidsAPIRequest) SetListAccountIds(_listAccountIds *ListAccountsByAccountIdsRequest) error {
	r._listAccountIds = _listAccountIds
	r.Set("list_account_ids", _listAccountIds)
	return nil
}

// Get ListAccountIds Getter
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetListAccountIds() *ListAccountsByAccountIdsRequest {
	return r._listAccountIds
}
