package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalLogisticsCarrierQuerylist 实际承运商查询
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
func CainiaoGlobalLogisticsCarrierQuerylist(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalLogisticsCarrierQuerylistAPIRequest, resp *cainiaohandover.CainiaoGlobalLogisticsCarrierQuerylistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
