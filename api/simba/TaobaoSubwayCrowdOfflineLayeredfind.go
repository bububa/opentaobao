package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCrowdOfflineLayeredfind 获取人群离线报表30转化周期
// taobao.subway.crowd.offline.layeredfind
//
// 获取人群离线报表
func TaobaoSubwayCrowdOfflineLayeredfind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayCrowdOfflineLayeredfindAPIRequest, resp *simba.TaobaoSubwayCrowdOfflineLayeredfindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
