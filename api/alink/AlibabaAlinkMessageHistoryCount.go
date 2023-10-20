package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkmessagehistorycount 查询消息总数
// alibaba.alink.message.history.count
//
// 查询消息总数
func Alibabaalinkmessagehistorycount(clt *core.SDKClient, req *alink.AlibabaalinkmessagehistorycountAPIRequest, session string) (*alink.AlibabaalinkmessagehistorycountAPIResponse, error) {
	var resp alink.AlibabaalinkmessagehistorycountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
