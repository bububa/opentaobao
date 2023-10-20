package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCrowdFindrecommendcrowd 查询推荐人群
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
func TaobaoUniversalbpCrowdFindrecommendcrowd(clt *core.SDKClient, req *simba.TaobaoUniversalbpCrowdFindrecommendcrowdAPIRequest, resp *simba.TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
