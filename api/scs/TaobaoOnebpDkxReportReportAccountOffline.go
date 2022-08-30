package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportAccountOffline 获取账户历史报表
// taobao.onebp.dkx.report.report.account.offline
//
// 获取账户历史报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{     "start_time": "2021-07-24",     "effect": 15,     "end_time": "2021-08-21",     "strategy_scene":true,     "unify_type":"kuan",     "bizCode":"adStrategyDkx"   }
func TaobaoOnebpDkxReportReportAccountOffline(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportAccountOfflineAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportAccountOfflineAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportAccountOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
