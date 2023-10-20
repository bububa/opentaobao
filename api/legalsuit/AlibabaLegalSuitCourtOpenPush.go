package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcourtopenpush 开庭信息推送接口
// alibaba.legal.suit.court.open.push
//
// 供ISV推送开庭信息给集团诉讼
func Alibabalegalsuitcourtopenpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcourtopenpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitcourtopenpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcourtopenpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
