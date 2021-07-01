package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaeBookBillGetAPIRequest
tae查询单笔虚拟账户明细 API请求
taobao.tae.book.bill.get

tae查询单笔虚拟账户明细 */
type TaobaoTaeBookBillGetAPIRequest struct {
	model.Params
	// 虚拟账户流水编号
	_bid int64
	// 需要返回的字段:参见BookBill结构体
	_fields []string
	// 虚拟账户流水编号
	_id int64
	// 虚拟账户科目编号
	_accountId int64
}

// New
