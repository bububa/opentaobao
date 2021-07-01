package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoElifeLifecardRefundAPIRequest
品牌惠卡券冲正退还 API请求
taobao.elife.lifecard.refund

淘宝生活汇消费卡虚拟卡，线下冲正退货接口 */
type TaobaoElifeLifecardRefundAPIRequest struct {
	model.Params
	// 请求参数
	_refundRequest *RefundRequest
}

// New
