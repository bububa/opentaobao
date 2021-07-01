package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountDeleteAPIRequest
OpenAccount删除数据 API请求
taobao.open.account.delete

OpenAccount删除数据 */
type TaobaoOpenAccountDeleteAPIRequest struct {
	model.Params
	// Open Account的id列表
	_openAccountIds []int64
	// ISV自己账号的id列表
	_isvAccountIds []string
}

// New
