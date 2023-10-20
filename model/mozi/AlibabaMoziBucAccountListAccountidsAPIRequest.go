package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziBucAccountListAccountidsAPIRequest 根据一批账号ID查询账号列表 API请求
// alibaba.mozi.buc.account.list.accountids
//
// 根据一批账号ID查询账号列表
type AlibabaMoziBucAccountListAccountidsAPIRequest struct {
	model.Params
	// 请求参数
	_listAccountIds *ListAccountsByAccountIdsRequest
}

// NewAlibabaMoziBucAccountListAccountidsRequest 初始化AlibabaMoziBucAccountListAccountidsAPIRequest对象
func NewAlibabaMoziBucAccountListAccountidsRequest() *AlibabaMoziBucAccountListAccountidsAPIRequest {
	return &AlibabaMoziBucAccountListAccountidsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziBucAccountListAccountidsAPIRequest) Reset() {
	r._listAccountIds = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.buc.account.list.accountids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetListAccountIds is ListAccountIds Setter
// 请求参数
func (r *AlibabaMoziBucAccountListAccountidsAPIRequest) SetListAccountIds(_listAccountIds *ListAccountsByAccountIdsRequest) error {
	r._listAccountIds = _listAccountIds
	r.Set("list_account_ids", _listAccountIds)
	return nil
}

// GetListAccountIds ListAccountIds Getter
func (r AlibabaMoziBucAccountListAccountidsAPIRequest) GetListAccountIds() *ListAccountsByAccountIdsRequest {
	return r._listAccountIds
}

var poolAlibabaMoziBucAccountListAccountidsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziBucAccountListAccountidsRequest()
	},
}

// GetAlibabaMoziBucAccountListAccountidsRequest 从 sync.Pool 获取 AlibabaMoziBucAccountListAccountidsAPIRequest
func GetAlibabaMoziBucAccountListAccountidsAPIRequest() *AlibabaMoziBucAccountListAccountidsAPIRequest {
	return poolAlibabaMoziBucAccountListAccountidsAPIRequest.Get().(*AlibabaMoziBucAccountListAccountidsAPIRequest)
}

// ReleaseAlibabaMoziBucAccountListAccountidsAPIRequest 将 AlibabaMoziBucAccountListAccountidsAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziBucAccountListAccountidsAPIRequest(v *AlibabaMoziBucAccountListAccountidsAPIRequest) {
	v.Reset()
	poolAlibabaMoziBucAccountListAccountidsAPIRequest.Put(v)
}
