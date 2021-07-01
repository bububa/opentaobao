package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaSalestarAdgroupFindbycampid
(销量明星)批量获取推广计划下的推广组信息
taobao.simba.salestar.adgroup.findbycampid

批量得到推广计划下的推广组 */
func TaobaoSimbaSalestarAdgroupFindbycampid(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest, session string) (*simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse, error) {
	var resp simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
