package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherCancelvoucherAPIRequest
作废券 API请求
taobao.game.deliveryvoucher.cancelvoucher

提货券发券接口：同步券和订单的关联信息 */
type TaobaoGameDeliveryvoucherCancelvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *CancelVoucherRequest
}

// New
