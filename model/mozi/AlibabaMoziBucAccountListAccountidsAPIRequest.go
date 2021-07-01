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

// New
