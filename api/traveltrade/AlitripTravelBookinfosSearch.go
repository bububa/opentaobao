package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelBookinfosSearch 飞猪度假-订单预定信息列表搜索接口
// alitrip.travel.bookinfos.search
//
// 查询订单预定信息列表
func AlitripTravelBookinfosSearch(clt *core.SDKClient, req *traveltrade.AlitripTravelBookinfosSearchAPIRequest, resp *traveltrade.AlitripTravelBookinfosSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
