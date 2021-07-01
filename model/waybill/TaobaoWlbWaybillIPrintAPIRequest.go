package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillIPrintAPIRequest
打印确认接口v1.0 API请求
taobao.wlb.waybill.i.print

打印面单前的校验接口，判断面单号信息与订单信息是否匹配。 */
type TaobaoWlbWaybillIPrintAPIRequest struct {
	model.Params
	// 打印请求
	_waybillApplyPrintCheckRequest *WaybillApplyPrintCheckRequest
}

// New
