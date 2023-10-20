package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabafcmallxinteractionaipiclist 花园ai作画定制查询
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
func Alibabafcmallxinteractionaipiclist(clt *core.SDKClient, req *interact.AlibabafcmallxinteractionaipiclistAPIRequest, session string) (*interact.AlibabafcmallxinteractionaipiclistAPIResponse, error) {
	var resp interact.AlibabafcmallxinteractionaipiclistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
