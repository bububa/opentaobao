package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryAccount 账户报表查询
// taobao.universalbp.report.query.account
//
// 账户报表查询
func TaobaoUniversalbpReportQueryAccount(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAccountAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryAccountAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryAccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
