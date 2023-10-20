package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalcommithandovercontentupdate 修改已经提交的交接单
// cainiao.global.commithandovercontent.update
//
// 修改已经提交的交接单
func Cainiaoglobalcommithandovercontentupdate(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalcommithandovercontentupdateAPIRequest, session string) (*cainiaohandover.CainiaoglobalcommithandovercontentupdateAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalcommithandovercontentupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
