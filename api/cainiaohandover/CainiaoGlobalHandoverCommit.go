package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

/* CainiaoGlobalHandoverCommit
提交发布交接单
cainiao.global.handover.commit

提供给ISV通过该接口提交发布交接单 */
func CainiaoGlobalHandoverCommit(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverCommitAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverCommitAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalHandoverCommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
