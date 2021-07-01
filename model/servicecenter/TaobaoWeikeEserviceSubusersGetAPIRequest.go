package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikeEserviceSubusersGetAPIRequest
客服外包订单分配的商家子账号列表 API请求
taobao.weike.eservice.subusers.get

获取客服外包订单分配的商家子账号列表，以及授权状态 */
type TaobaoWeikeEserviceSubusersGetAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// New
