package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBillBookBillsGetAPIRequest
查询虚拟账户明细数据(自研发商家专用) API请求
taobao.bill.book.bills.get

查询虚拟账户明细数据 */
type TaobaoBillBookBillsGetAPIRequest struct {
	model.Params
	// 虚拟账户科目编号
	_accountId int64
	// 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
	_journalTypes []int64
	// 记账开始时间
	_startTime string
	// 记账结束时间,end_time与start_time相差不能超过30天
	_endTime string
	// 每页大小,建议40~100,不能超过100
	_pageSize int64
	// 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
	_fields string
	// 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
	_pageNo int64
}

// New
