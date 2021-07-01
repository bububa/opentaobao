package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziBucAccountPageallAPIRequest
查询租户内内所有账号 API请求
alibaba.mozi.buc.account.pageall

查询租户内内所有账号 */
type AlibabaMoziBucAccountPageallAPIRequest struct {
	model.Params
	// 查询租户内所有人员和账号
	_pageAll *PageAllAccountsRequest
}

// New
