package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// Tmallfansarenapush 消息推送
// tmall.fans.arena.push
//
// 超级擂台消息推送
func Tmallfansarenapush(clt *core.SDKClient, req *fans.TmallfansarenapushAPIRequest, session string) (*fans.TmallfansarenapushAPIResponse, error) {
	var resp fans.TmallfansarenapushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
