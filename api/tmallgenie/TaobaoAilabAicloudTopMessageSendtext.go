package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmessagesendtext 故事机发送文本留言
// taobao.ailab.aicloud.top.message.sendtext
//
// 故事机文本留言
func Taobaoailabaicloudtopmessagesendtext(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmessagesendtextAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmessagesendtextAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmessagesendtextAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
