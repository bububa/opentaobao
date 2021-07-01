package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountListAPIRequest
OpenAccount账号信息查询 API请求
taobao.open.account.list

OpenAccount账号信息查询 */
type TaobaoOpenAccountListAPIRequest struct {
	model.Params
	// Open Account的id列表, 每次最多查询 20 个帐户
	_openAccountIds []int64
	// ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
	_isvAccountIds []string
}

// New
