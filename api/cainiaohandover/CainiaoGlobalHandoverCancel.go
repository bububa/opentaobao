package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandovercancel 取消交接单
// cainiao.global.handover.cancel
//
// 提供给ISV通过该接口取消交接单
func Cainiaoglobalhandovercancel(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandovercancelAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandovercancelAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandovercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
