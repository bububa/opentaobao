package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

/* CainiaoWaybillIiQueryByWaybillcode
通过面单号查询面单信息
cainiao.waybill.ii.query.by.waybillcode

通过面单号查看面单号的当前状态，如签收、发货、失效等。 */
func CainiaoWaybillIiQueryByWaybillcode(clt *core.SDKClient, req *waybill.CainiaoWaybillIiQueryByWaybillcodeAPIRequest, session string) (*waybill.CainiaoWaybillIiQueryByWaybillcodeAPIResponse, error) {
	var resp waybill.CainiaoWaybillIiQueryByWaybillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
