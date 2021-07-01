package deliveryvoucher

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherSendgoodsAPIRequest
提货券发货接口 API请求
taobao.game.deliveryvoucher.sendgoods

提货券发券接口：同步券和订单的关联信息 */
type TaobaoGameDeliveryvoucherSendgoodsAPIRequest struct {
	model.Params
	// 发券参数
	_param0 *SendVoucherRequest
}

// New
