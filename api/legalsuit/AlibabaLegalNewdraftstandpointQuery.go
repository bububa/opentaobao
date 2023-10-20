package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalNewdraftstandpointQuery 未采纳口径查询(新)
// alibaba.legal.newdraftstandpoint.query
//
// 未采纳口径查询(新)
func AlibabaLegalNewdraftstandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalNewdraftstandpointQueryAPIRequest, session string) (*legalsuit.AlibabaLegalNewdraftstandpointQueryAPIResponse, error) {
	var resp legalsuit.AlibabaLegalNewdraftstandpointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
