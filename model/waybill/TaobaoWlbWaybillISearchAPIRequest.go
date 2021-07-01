package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillISearchAPIRequest
查询面单服务订购及面单使用情况v1.0 API请求
taobao.wlb.waybill.i.search

获取发货地&CP开通状态&账户的使用情况 */
type TaobaoWlbWaybillISearchAPIRequest struct {
	model.Params
	// 查询网点信息
	_waybillApplyRequest *WaybillApplyRequest
}

// New
