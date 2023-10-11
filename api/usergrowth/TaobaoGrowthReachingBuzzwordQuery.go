package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingBuzzwordQuery 淘宝热词榜单数据查询接口
// taobao.growth.reaching.buzzword.query
//
// 查询淘宝热词榜单数据
func TaobaoGrowthReachingBuzzwordQuery(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIRequest, session string) (*usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIResponse, error) {
	var resp usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
