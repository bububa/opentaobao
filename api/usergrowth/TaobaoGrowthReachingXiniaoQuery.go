package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaogrowthreachingxiniaoquery 查询溪鸟推荐信息数据
// taobao.growth.reaching.xiniao.query
//
// 查询溪鸟推荐信息数据
func Taobaogrowthreachingxiniaoquery(clt *core.SDKClient, req *usergrowth.TaobaogrowthreachingxiniaoqueryAPIRequest, session string) (*usergrowth.TaobaogrowthreachingxiniaoqueryAPIResponse, error) {
	var resp usergrowth.TaobaogrowthreachingxiniaoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
