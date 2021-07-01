package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaRptAdgroupeffectGet
推广组效果报表数据对象
taobao.simba.rpt.adgroupeffect.get

推广组效果报表数据对象 */
func TaobaoSimbaRptAdgroupeffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupeffectGetAPIRequest, session string) (*simba.TaobaoSimbaRptAdgroupeffectGetAPIResponse, error) {
	var resp simba.TaobaoSimbaRptAdgroupeffectGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
