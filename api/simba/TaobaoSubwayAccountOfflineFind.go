package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAccountOfflineFind 获取账户历史多日汇总报表
// taobao.subway.account.offline.find
//
// 获取账户历史报表
func TaobaoSubwayAccountOfflineFind(clt *core.SDKClient, req *simba.TaobaoSubwayAccountOfflineFindAPIRequest, resp *simba.TaobaoSubwayAccountOfflineFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
