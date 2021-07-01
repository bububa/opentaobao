package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiGetAPIRequest
电子面单云打印接口 API请求
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法 */
type CainiaoWaybillIiGetAPIRequest struct {
	model.Params
	// 入参信息
	_paramWaybillCloudPrintApplyNewRequest *WaybillCloudPrintApplyNewRequest
}

// New
