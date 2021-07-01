package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsWaybillGetAPIRequest
获取运单信息 API请求
taobao.wlb.imports.waybill.get

一般进口商家，获取订单的电子面单链接地址 */
type TaobaoWlbImportsWaybillGetAPIRequest struct {
	model.Params
	// 物流订单号
	_orderCode string
}

// New
