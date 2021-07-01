package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderUpdateConfirmcodeAPIRequest
推送及更新订单确认号 API请求
taobao.xhotel.order.update.confirmcode

商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。 */
type TaobaoXhotelOrderUpdateConfirmcodeAPIRequest struct {
	model.Params
	// 系统入参
	_param *UpdateOrderConfirmCodeParam
}

// New
