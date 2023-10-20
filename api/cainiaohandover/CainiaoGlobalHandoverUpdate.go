package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverUpdate 修改交接单
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
func CainiaoGlobalHandoverUpdate(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverUpdateAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
