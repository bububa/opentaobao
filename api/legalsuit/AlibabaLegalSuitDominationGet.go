package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitDominationGet 查询管辖信息
// alibaba.legal.suit.domination.get
//
// 查询管辖信息
func AlibabaLegalSuitDominationGet(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitDominationGetAPIRequest, session string) (*legalsuit.AlibabaLegalSuitDominationGetAPIResponse, error) {
	var resp legalsuit.AlibabaLegalSuitDominationGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
