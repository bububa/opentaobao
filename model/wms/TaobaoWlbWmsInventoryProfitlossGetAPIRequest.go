package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsInventoryProfitlossGetAPIRequest
通过订单列表批量获取库存损益单信息 API请求
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息 */
type TaobaoWlbWmsInventoryProfitlossGetAPIRequest struct {
	model.Params
	// 菜鸟订单编码
	_cnOrderCode string
}

// New
