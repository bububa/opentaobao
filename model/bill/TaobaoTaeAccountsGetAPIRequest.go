package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaeAccountsGetAPIRequest
tae查询费用科目信息 API请求
taobao.tae.accounts.get

tae查询费用科目信息 */
type TaobaoTaeAccountsGetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []int64
}

// New
