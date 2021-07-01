package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillICancelAPIRequest
商家取消获取的电子面单号v1.0 API请求
taobao.wlb.waybill.i.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。 */
type TaobaoWlbWaybillICancelAPIRequest struct {
	model.Params
	// 取消接口入参
	_waybillApplyCancelRequest *WaybillApplyCancelRequest
}

// New
