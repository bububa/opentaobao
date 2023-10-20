package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillAddressReachableQuery 地址可达查询
// cainiao.waybill.address.reachable.query
//
// 地址可达查询
func CainiaoWaybillAddressReachableQuery(clt *core.SDKClient, req *waybill.CainiaoWaybillAddressReachableQueryAPIRequest, resp *waybill.CainiaoWaybillAddressReachableQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
