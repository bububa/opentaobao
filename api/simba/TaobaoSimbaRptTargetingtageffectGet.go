package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptTargetingtageffectGet 获取定向效果报表数据
// taobao.simba.rpt.targetingtageffect.get
//
// 获取定向效果报表数据
func TaobaoSimbaRptTargetingtageffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtageffectGetAPIRequest, resp *simba.TaobaoSimbaRptTargetingtageffectGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
