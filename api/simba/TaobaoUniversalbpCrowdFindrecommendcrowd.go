package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcrowdfindrecommendcrowd 查询推荐人群
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
func Taobaouniversalbpcrowdfindrecommendcrowd(clt *core.SDKClient, req *simba.TaobaouniversalbpcrowdfindrecommendcrowdAPIRequest, session string) (*simba.TaobaouniversalbpcrowdfindrecommendcrowdAPIResponse, error) {
	var resp simba.TaobaouniversalbpcrowdfindrecommendcrowdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
