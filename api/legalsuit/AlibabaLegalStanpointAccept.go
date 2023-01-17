package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStanpointAccept 采纳口径
// alibaba.legal.stanpoint.accept
//
// 采纳口径
func AlibabaLegalStanpointAccept(clt *core.SDKClient, req *legalsuit.AlibabaLegalStanpointAcceptAPIRequest, session string) (*legalsuit.AlibabaLegalStanpointAcceptAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStanpointAcceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
