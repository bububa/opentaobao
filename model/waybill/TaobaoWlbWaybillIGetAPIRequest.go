package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillIGetAPIRequest
获取物流服务商电子面单号v1.0 API请求
taobao.wlb.waybill.i.get

商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。 */
type TaobaoWlbWaybillIGetAPIRequest struct {
	model.Params
	// 面单申请
	_waybillApplyNewRequest *WaybillApplyNewRequest
}

// New
