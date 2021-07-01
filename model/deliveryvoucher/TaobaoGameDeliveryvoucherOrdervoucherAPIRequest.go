package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherOrdervoucherAPIRequest
预约接口 API请求
taobao.game.deliveryvoucher.ordervoucher

提货券发券接口：同步券和订单的关联信息 */
type TaobaoGameDeliveryvoucherOrdervoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *OrderVoucherRequest
}

// New
