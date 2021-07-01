package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSpecialRefundGetAPIRequest
特殊部分退纠纷单查询 API请求
taobao.special.refund.get

获取单笔特殊部分退的纠纷单查询 */
type TaobaoSpecialRefundGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
	_fields []string
	// 退款单号
	_refundId int64
}

// New
