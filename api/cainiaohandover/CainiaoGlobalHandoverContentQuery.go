package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverContentQuery 查询大包详情
// cainiao.global.handover.content.query
//
// 查询大包详情
func CainiaoGlobalHandoverContentQuery(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverContentQueryAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverContentQueryAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalHandoverContentQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
