package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalCommithandovercontentUpdate 修改已经提交的交接单
// cainiao.global.commithandovercontent.update
//
// 修改已经提交的交接单
func CainiaoGlobalCommithandovercontentUpdate(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIRequest, session string) (*cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
