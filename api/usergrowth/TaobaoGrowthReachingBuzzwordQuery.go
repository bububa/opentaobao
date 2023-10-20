package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingBuzzwordQuery 淘宝热词榜单数据查询接口
// taobao.growth.reaching.buzzword.query
//
// 查询淘宝热词榜单数据
func TaobaoGrowthReachingBuzzwordQuery(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIRequest, resp *usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
