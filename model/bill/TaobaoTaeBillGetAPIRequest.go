package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaeBillGetAPIRequest
tae查询单笔账单明细 API请求
taobao.tae.bill.get

查询单笔账单明细 */
type TaobaoTaeBillGetAPIRequest struct {
	model.Params
	// 账单编号
	_bid int64
	// 传入需要返回的字段
	_fields []string
	// 账单编号
	_id int64
	// 虚拟账户科目编号
	_accountId int64
}

// New
