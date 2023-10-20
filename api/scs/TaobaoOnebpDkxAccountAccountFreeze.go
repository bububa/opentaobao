package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxaccountaccountfreeze 创建计划后支付
// taobao.onebp.dkx.account.account.freeze
//
// 创建计划后支付。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func Taobaoonebpdkxaccountaccountfreeze(clt *core.SDKClient, req *scs.TaobaoonebpdkxaccountaccountfreezeAPIRequest, session string) (*scs.TaobaoonebpdkxaccountaccountfreezeAPIResponse, error) {
	var resp scs.TaobaoonebpdkxaccountaccountfreezeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
