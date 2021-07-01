package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaeBillsGetAPIRequest
tae查询账单明细 API请求
taobao.tae.bills.get

tae查询账单明细 */
type TaobaoTaeBillsGetAPIRequest struct {
	model.Params
	// 开始时间
	_queryStartDate string
	// 交易编号
	_pTradeId int64
	// 子订单编号
	_tradeId int64
	// 每页大小,默认40条,可选范围 ：40~100
	_pageSize int64
	// 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序
	_queryDateType int64
	// 页数,建议不要超过100页,越大性能越低,有可能会超时
	_currentPage int64
	// 科目编号
	_itemId int64
	// 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外)
	_queryEndDate string
	// 传入需要返回的字段,参见Bill结构体
	_fields []string
}

// New
