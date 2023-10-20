package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalLogisticsCarrierQuerylist 实际承运商查询
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
func CainiaoGlobalLogisticsCarrierQuerylist(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalLogisticsCarrierQuerylistAPIRequest, session string) (*cainiaohandover.CainiaoGlobalLogisticsCarrierQuerylistAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalLogisticsCarrierQuerylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
