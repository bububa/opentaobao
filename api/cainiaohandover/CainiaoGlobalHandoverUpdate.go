package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandoverupdate 修改交接单
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
func Cainiaoglobalhandoverupdate(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandoverupdateAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandoverupdateAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandoverupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
