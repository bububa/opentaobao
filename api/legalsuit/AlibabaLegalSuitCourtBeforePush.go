package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcourtbeforepush 更新或保存庭前信息
// alibaba.legal.suit.court.before.push
//
// 更新或者保存庭前信息
func Alibabalegalsuitcourtbeforepush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcourtbeforepushAPIRequest, session string) (*legalsuit.AlibabalegalsuitcourtbeforepushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcourtbeforepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
