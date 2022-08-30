package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportAccountRealtime 获取账户实时报表
// taobao.onebp.dkx.report.report.account.realtime
//
// 获取账户实时报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{     "log_date_list": [  "2021-09-23"     ]   }
func TaobaoOnebpDkxReportReportAccountRealtime(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportAccountRealtimeAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportAccountRealtimeAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportAccountRealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
