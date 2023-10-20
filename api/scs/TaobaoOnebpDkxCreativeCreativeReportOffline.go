package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCreativeCreativeReportOffline 获取创意离线报表
// taobao.onebp.dkx.creative.creative.report.offline
//
// 获取创意离线报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCreativeCreativeReportOffline(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest, resp *scs.TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
