package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgTpwdRiskReport 淘宝客-推广者-淘口令预警及拦截查询
// taobao.tbk.dg.tpwd.risk.report
//
// 淘宝客-推广者-淘口令预警及拦截查询
func TaobaoTbkDgTpwdRiskReport(clt *core.SDKClient, req *tbk.TaobaoTbkDgTpwdRiskReportAPIRequest, session string) (*tbk.TaobaoTbkDgTpwdRiskReportAPIResponse, error) {
	var resp tbk.TaobaoTbkDgTpwdRiskReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
