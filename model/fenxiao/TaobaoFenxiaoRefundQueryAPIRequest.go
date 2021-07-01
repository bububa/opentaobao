package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoRefundQueryAPIRequest
批量查询采购退款 API请求
taobao.fenxiao.refund.query

供应商按查询条件批量查询代销采购退款 */
type TaobaoFenxiaoRefundQueryAPIRequest struct {
	model.Params
	// 代销采购退款单最早修改时间
	_startDate string
	// 代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天
	_endDate string
	// 页码（大于0的整数。无值或小于1的值按默认值1计）
	_pageNo int64
	// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
	_pageSize int64
	// 是否查询下游买家的退款信息
	_querySellerRefund bool
	// 渠道code，可批量 老供销渠道：999
	_tradeTypes []int64
	// 角色，供应商：2，分销商：1
	_userRoleType int64
	// 代销：1 经销：2 寄售（猫超自营寄售）：5 平台寄售：6
	_channelCodes []int64
}

// New
