package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiUpdateAPIRequest
电子面单云打印更新接口 API请求
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。 */
type CainiaoWaybillIiUpdateAPIRequest struct {
	model.Params
	// 更新请求信息
	_paramWaybillCloudPrintUpdateRequest *WaybillCloudPrintUpdateRequest
}

// New
