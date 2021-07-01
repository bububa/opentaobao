package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherEvaluateAPIRequest
卡券评价回传 API请求
taobao.game.deliveryvoucher.evaluate

卡券ISV回传商品评价 */
type TaobaoGameDeliveryvoucherEvaluateAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *VoucherEvaluateRequest
}

// New
