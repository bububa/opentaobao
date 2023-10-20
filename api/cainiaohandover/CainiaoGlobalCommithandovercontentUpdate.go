package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalCommithandovercontentUpdate 修改已经提交的交接单
// cainiao.global.commithandovercontent.update
//
// 修改已经提交的交接单
func CainiaoGlobalCommithandovercontentUpdate(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIRequest, resp *cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
