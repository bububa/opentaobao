package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandovercontentsubbagadd 预约单下追加大包
// cainiao.global.handover.content.subbag.add
//
// 预约单下追加大包
func Cainiaoglobalhandovercontentsubbagadd(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandovercontentsubbagaddAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandovercontentsubbagaddAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandovercontentsubbagaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
