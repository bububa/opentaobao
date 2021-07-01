package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRefundRefuseAPIRequest
卖家拒绝退款 API请求
taobao.refund.refuse

卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：<br/>1. 传入的refund_id和相应的tid, oid必须匹配<br/>2. 如果一笔订单只有一笔子订单，则tid必须与oid相同<br/>3. 只有卖家才能执行拒绝退款操作<br/>4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单 */
type TaobaoRefundRefuseAPIRequest struct {
	model.Params
	// 退款单号
	_refundId int64
	// 拒绝退款时的说明信息，长度2-200
	_refuseMessage string
	// 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。
	_refuseProof *model.File
	// 可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。
	_refundPhase string
	// 退款版本号，天猫退款为必填项。
	_refundVersion int64
	// 拒绝原因编号，会提供用户拒绝原因列表供选择
	_refuseReasonId int64
}

// New
