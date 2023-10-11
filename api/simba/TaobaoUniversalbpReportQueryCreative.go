package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCreative 创意报表查询
// taobao.universalbp.report.query.creative
//
// 创意报表查询
func TaobaoUniversalbpReportQueryCreative(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCreativeAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryCreativeAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryCreativeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
