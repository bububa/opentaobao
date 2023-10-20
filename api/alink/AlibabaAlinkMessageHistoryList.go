package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Alibabaalinkmessagehistorylist 查询消息列表
// alibaba.alink.message.history.list
//
// 查询消息列表
func Alibabaalinkmessagehistorylist(clt *core.SDKClient, req *alink.AlibabaalinkmessagehistorylistAPIRequest, session string) (*alink.AlibabaalinkmessagehistorylistAPIResponse, error) {
	var resp alink.AlibabaalinkmessagehistorylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
