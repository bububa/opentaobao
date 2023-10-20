package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverContentQuery 查询大包详情
// cainiao.global.handover.content.query
//
// 查询大包详情
func CainiaoGlobalHandoverContentQuery(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverContentQueryAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverContentQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
