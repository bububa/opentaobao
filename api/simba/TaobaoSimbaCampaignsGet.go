package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaCampaignsGet
取得一组推广计划
taobao.simba.campaigns.get

取得一个客户的推广计划； */
func TaobaoSimbaCampaignsGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignsGetAPIRequest, session string) (*simba.TaobaoSimbaCampaignsGetAPIResponse, error) {
	var resp simba.TaobaoSimbaCampaignsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
