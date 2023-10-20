package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiQueryByWaybillcode 通过面单号查询面单打印报文
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查询面单的打印报文
func CainiaoWaybillIiQueryByWaybillcode(clt *core.SDKClient, req *waybill.CainiaoWaybillIiQueryByWaybillcodeAPIRequest, session string) (*waybill.CainiaoWaybillIiQueryByWaybillcodeAPIResponse, error) {
	var resp waybill.CainiaoWaybillIiQueryByWaybillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
