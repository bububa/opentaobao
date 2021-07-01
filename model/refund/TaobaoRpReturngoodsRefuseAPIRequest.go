package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRpReturngoodsRefuseAPIRequest
卖家拒绝退货 API请求
taobao.rp.returngoods.refuse

卖家拒绝退货，目前仅支持天猫退货。 */
type TaobaoRpReturngoodsRefuseAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 退款服务状态，售后或者售中
	_refundPhase string
	// 退款版本号
	_refundVersion int64
	// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
	_refuseProof *model.File
	// 拒绝原因编号，会提供拒绝原因列表供选择
	_refuseReasonId int64
}

// New
