package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAccountOfflineLayeredfind 获取账户历史报表30天转化周期
// taobao.subway.account.offline.layeredfind
//
// 获取账户历史报表
func TaobaoSubwayAccountOfflineLayeredfind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayAccountOfflineLayeredfindAPIRequest, resp *simba.TaobaoSubwayAccountOfflineLayeredfindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
