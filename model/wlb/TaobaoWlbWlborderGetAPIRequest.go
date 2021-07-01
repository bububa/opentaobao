package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWlborderGetAPIRequest
根据物流宝订单编号查询物流宝订单概要信息 API请求
taobao.wlb.wlborder.get

根据物流宝订单编号查询物流宝订单概要信息 */
type TaobaoWlbWlborderGetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_wlbOrderCode string
}

// New
