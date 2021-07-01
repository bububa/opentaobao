package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaRptTargetingtagGet
搜索人群离线报表
taobao.simba.rpt.targetingtag.get

获取搜搜人群实时报表 */
func TaobaoSimbaRptTargetingtagGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtagGetAPIRequest, session string) (*simba.TaobaoSimbaRptTargetingtagGetAPIResponse, error) {
	var resp simba.TaobaoSimbaRptTargetingtagGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
