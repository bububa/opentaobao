package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStanpointAccept 采纳口径
// alibaba.legal.stanpoint.accept
//
// 采纳口径
func AlibabaLegalStanpointAccept(clt *core.SDKClient, req *legalsuit.AlibabaLegalStanpointAcceptAPIRequest, resp *legalsuit.AlibabaLegalStanpointAcceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
