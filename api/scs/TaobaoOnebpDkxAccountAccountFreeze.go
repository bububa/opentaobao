package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxAccountAccountFreeze 创建计划后支付
// taobao.onebp.dkx.account.account.freeze
//
// 创建计划后支付。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxAccountAccountFreeze(clt *core.SDKClient, req *scs.TaobaoOnebpDkxAccountAccountFreezeAPIRequest, resp *scs.TaobaoOnebpDkxAccountAccountFreezeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
