package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBillAccountsGetAPIRequest
查询费用科目信息(限自研商家) API请求
taobao.bill.accounts.get

查询费用账户信息 */
type TaobaoBillAccountsGetAPIRequest struct {
	model.Params
	// 需要返回的字段
	_fields []string
	// 需要获取的科目ID
	_aids []int64
}

// New
