package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgTpwdReportGet 淘宝客-推广者-淘口令回流数据查询
// taobao.tbk.dg.tpwd.report.get
//
// 淘宝客获取单个淘口令的回流PV、UV数据。
func TaobaoTbkDgTpwdReportGet(clt *core.SDKClient, req *tbk.TaobaoTbkDgTpwdReportGetAPIRequest, session string) (*tbk.TaobaoTbkDgTpwdReportGetAPIResponse, error) {
	var resp tbk.TaobaoTbkDgTpwdReportGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
