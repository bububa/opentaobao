package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitSealPush 法宝推送用印
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
func AlibabaLegalSuitSealPush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitSealPushAPIRequest, session string) (*legalsuit.AlibabaLegalSuitSealPushAPIResponse, error) {
	var resp legalsuit.AlibabaLegalSuitSealPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
