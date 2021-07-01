package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillIProductAPIRequest
商家查询物流商产品类型接口 API请求
taobao.wlb.waybill.i.product

商家可以查询物流商的产品类型和服务能力。 */
type TaobaoWlbWaybillIProductAPIRequest struct {
	model.Params
	// 查询物流商电子面单产品类型入参
	_waybillProductTypeRequest *WaybillProductTypeRequest
}

// New
