package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundMessageGetAPIRequest
openmall获取退款单留言 API请求
taobao.openmall.refund.message.get

openmall获取退款单留言 */
type TaobaoOpenmallRefundMessageGetAPIRequest struct {
	model.Params
	// 分销者身份
	_distributor string
	// 翻页页码
	_pageNo int64
	// 翻页大小
	_pageSize int64
	// 退款单号
	_refundId int64
}

// New
