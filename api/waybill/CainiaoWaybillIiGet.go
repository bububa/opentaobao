package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiGet 电子面单云打印接口
// cainiao.waybill.ii.get
//
// 菜鸟电子面单的云打印申请电子面单号的方法
func CainiaoWaybillIiGet(clt *core.SDKClient, req *waybill.CainiaoWaybillIiGetAPIRequest, resp *waybill.CainiaoWaybillIiGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
