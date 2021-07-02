package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverUpdate 修改交接单
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
func CainiaoGlobalHandoverUpdate(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverUpdateAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverUpdateAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalHandoverUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
