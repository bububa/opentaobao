package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcourtlawyerpush 推荐律师接口
// alibaba.legal.suit.court.lawyer.push
//
// 为诉讼系统推荐律师
func Alibabalegalsuitcourtlawyerpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcourtlawyerpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitcourtlawyerpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcourtlawyerpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
