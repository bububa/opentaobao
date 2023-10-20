package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliiquerybywaybillcode 通过面单号查询面单打印报文
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查询面单的打印报文
func Cainiaowaybilliiquerybywaybillcode(clt *core.SDKClient, req *waybill.CainiaowaybilliiquerybywaybillcodeAPIRequest, session string) (*waybill.CainiaowaybilliiquerybywaybillcodeAPIResponse, error) {
	var resp waybill.CainiaowaybilliiquerybywaybillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
