package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderstatusGetAPIRequest
物流宝订单流转状态查询 API请求
taobao.wlb.orderstatus.get

根据物流宝订单号查询物流宝订单至目前为止的流转状态列表 */
type TaobaoWlbOrderstatusGetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_orderCode string
}

// New
