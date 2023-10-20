package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandoversavedraft 创建交接单草稿
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
func Cainiaoglobalhandoversavedraft(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandoversavedraftAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandoversavedraftAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandoversavedraftAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
