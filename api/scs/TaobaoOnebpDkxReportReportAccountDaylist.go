package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportAccountDaylist 获取账户分日报表
// taobao.onebp.dkx.report.report.account.daylist
//
// 获取账户分日报表
func TaobaoOnebpDkxReportReportAccountDaylist(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportAccountDaylistAPIRequest, resp *scs.TaobaoOnebpDkxReportReportAccountDaylistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
