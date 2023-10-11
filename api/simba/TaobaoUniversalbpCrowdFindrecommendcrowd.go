package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCrowdFindrecommendcrowd 查询推荐人群
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
func TaobaoUniversalbpCrowdFindrecommendcrowd(clt *core.SDKClient, req *simba.TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest, session string) (*simba.TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse, error) {
	var resp simba.TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
