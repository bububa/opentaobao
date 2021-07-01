package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

/* TaobaoSimbaRptTargetingtagbaseGet
定向基础报表
taobao.simba.rpt.targetingtagbase.get

获取定向基础报表 */
func TaobaoSimbaRptTargetingtagbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtagbaseGetAPIRequest, session string) (*simba.TaobaoSimbaRptTargetingtagbaseGetAPIResponse, error) {
	var resp simba.TaobaoSimbaRptTargetingtagbaseGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
