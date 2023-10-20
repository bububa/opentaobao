package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiQueryByTradecode 通过订单号查询电子面单通接口
// cainiao.waybill.ii.query.by.tradecode
//
// 通过订单号查看面单的信息
func CainiaoWaybillIiQueryByTradecode(clt *core.SDKClient, req *waybill.CainiaoWaybillIiQueryByTradecodeAPIRequest, resp *waybill.CainiaoWaybillIiQueryByTradecodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
