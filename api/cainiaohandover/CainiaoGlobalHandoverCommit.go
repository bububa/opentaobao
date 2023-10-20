package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandovercommit 提交发布交接单
// cainiao.global.handover.commit
//
// 提供给ISV通过该接口提交发布交接单
func Cainiaoglobalhandovercommit(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandovercommitAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandovercommitAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandovercommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
