package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRefundMessagesGetAPIRequest
查询退款留言/凭证列表 API请求
taobao.refund.messages.get

查询退款留言/凭证列表 */
type TaobaoRefundMessagesGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
	_fields []string
	// 退款单号
	_refundId int64
	// 页码
	_pageNo int64
	// 每页条数
	_pageSize int64
	// 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
	_refundPhase string
}

// New
