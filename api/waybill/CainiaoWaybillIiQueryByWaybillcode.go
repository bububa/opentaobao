package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiQueryByWaybillcode 通过面单号查询面单打印报文
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查询面单的打印报文
func CainiaoWaybillIiQueryByWaybillcode(clt *core.SDKClient, req *waybill.CainiaoWaybillIiQueryByWaybillcodeAPIRequest, resp *waybill.CainiaoWaybillIiQueryByWaybillcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
