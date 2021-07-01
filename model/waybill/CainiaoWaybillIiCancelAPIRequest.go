package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiCancelAPIRequest
商家取消获取的电子面单号 API请求
cainiao.waybill.ii.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。 */
type CainiaoWaybillIiCancelAPIRequest struct {
	model.Params
	// 快递公司code
	_cpCode string
	// 电子面单号
	_waybillCode string
}

// New
