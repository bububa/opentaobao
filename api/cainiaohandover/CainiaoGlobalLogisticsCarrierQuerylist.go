package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaogloballogisticscarrierquerylist 实际承运商查询
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
func Cainiaogloballogisticscarrierquerylist(clt *core.SDKClient, req *cainiaohandover.CainiaogloballogisticscarrierquerylistAPIRequest, session string) (*cainiaohandover.CainiaogloballogisticscarrierquerylistAPIResponse, error) {
	var resp cainiaohandover.CainiaogloballogisticscarrierquerylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
