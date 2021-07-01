package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* TaobaoWeitaoFeedCancel
取消广播在timeline、广场中展示
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。 */
func TaobaoWeitaoFeedCancel(clt *core.SDKClient, req *interact.TaobaoWeitaoFeedCancelAPIRequest, session string) (*interact.TaobaoWeitaoFeedCancelAPIResponse, error) {
	var resp interact.TaobaoWeitaoFeedCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
