package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingXiniaoQuery 查询溪鸟推荐信息数据
// taobao.growth.reaching.xiniao.query
//
// 查询溪鸟推荐信息数据
func TaobaoGrowthReachingXiniaoQuery(clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingXiniaoQueryAPIRequest, session string) (*usergrowth.TaobaoGrowthReachingXiniaoQueryAPIResponse, error) {
	var resp usergrowth.TaobaoGrowthReachingXiniaoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
