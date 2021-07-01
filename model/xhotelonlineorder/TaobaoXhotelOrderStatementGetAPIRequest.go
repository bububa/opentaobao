package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderStatementGetAPIRequest
查询账单信息 API请求
taobao.xhotel.order.statement.get

阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情 */
type TaobaoXhotelOrderStatementGetAPIRequest struct {
	model.Params
	// 要查询的tid列表，逗号分隔,列表查询;当此值不为空时候，其余参数忽略。最多单次20条。
	_orderTids string
	// 查询条数，最大支持500条
	_pageSize int64
	// 数据查询开始下标
	_start int64
	// 0：check_in, 1：check_out,2：分账时间
	_dateType int64
	// 查询结束时间
	_to string
	// 查询开始时间
	_from string
	// 淘宝订单号
	_tid int64
	// 外部酒店编码
	_hotelCode string
	// 系统商vendor
	_vendor string
}

// New
