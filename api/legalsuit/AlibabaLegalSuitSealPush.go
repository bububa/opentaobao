package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitsealpush 法宝推送用印
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
func Alibabalegalsuitsealpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitsealpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitsealpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitsealpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
