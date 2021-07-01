package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherSendvoucherAPIRequest
提货券发券接口 API请求
taobao.game.deliveryvoucher.sendvoucher

提货券发券接口：同步券和订单的关联信息 */
type TaobaoGameDeliveryvoucherSendvoucherAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// New
