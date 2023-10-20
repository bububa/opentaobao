package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilladdressreachablequery 地址可达查询
// cainiao.waybill.address.reachable.query
//
// 地址可达查询
func Cainiaowaybilladdressreachablequery(clt *core.SDKClient, req *waybill.CainiaowaybilladdressreachablequeryAPIRequest, session string) (*waybill.CainiaowaybilladdressreachablequeryAPIResponse, error) {
	var resp waybill.CainiaowaybilladdressreachablequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
