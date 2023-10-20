package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverContentSubbagAdd 预约单下追加大包
// cainiao.global.handover.content.subbag.add
//
// 预约单下追加大包
func CainiaoGlobalHandoverContentSubbagAdd(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverContentSubbagAddAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverContentSubbagAddAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalHandoverContentSubbagAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
