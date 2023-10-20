package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalstanpointaccept 采纳口径
// alibaba.legal.stanpoint.accept
//
// 采纳口径
func Alibabalegalstanpointaccept(clt *core.SDKClient, req *legalsuit.AlibabalegalstanpointacceptAPIRequest, session string) (*legalsuit.AlibabalegalstanpointacceptAPIResponse, error) {
	var resp legalsuit.AlibabalegalstanpointacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
