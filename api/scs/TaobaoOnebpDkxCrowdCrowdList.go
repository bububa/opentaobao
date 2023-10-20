package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxcrowdcrowdlist 获取人群信息列表
// taobao.onebp.dkx.crowd.crowd.list
//
// 获取人群信息列表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func Taobaoonebpdkxcrowdcrowdlist(clt *core.SDKClient, req *scs.TaobaoonebpdkxcrowdcrowdlistAPIRequest, session string) (*scs.TaobaoonebpdkxcrowdcrowdlistAPIResponse, error) {
	var resp scs.TaobaoonebpdkxcrowdcrowdlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
