package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverSavedraft 创建交接单草稿
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
func CainiaoGlobalHandoverSavedraft(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverSavedraftAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverSavedraftAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
