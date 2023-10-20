package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Taobaoweitaofeedsynchronize 推广淘小铺isv 活动到微淘feed
// taobao.weitao.feed.synchronize
//
// 推广淘小铺isv 活动到微淘feed
func Taobaoweitaofeedsynchronize(clt *core.SDKClient, req *interact.TaobaoweitaofeedsynchronizeAPIRequest, session string) (*interact.TaobaoweitaofeedsynchronizeAPIResponse, error) {
	var resp interact.TaobaoweitaofeedsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
