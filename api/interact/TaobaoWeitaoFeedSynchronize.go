package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// TaobaoWeitaoFeedSynchronize 推广淘小铺isv 活动到微淘feed
// taobao.weitao.feed.synchronize
//
// 推广淘小铺isv 活动到微淘feed
func TaobaoWeitaoFeedSynchronize(clt *core.SDKClient, req *interact.TaobaoWeitaoFeedSynchronizeAPIRequest, session string) (*interact.TaobaoWeitaoFeedSynchronizeAPIResponse, error) {
	var resp interact.TaobaoWeitaoFeedSynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
