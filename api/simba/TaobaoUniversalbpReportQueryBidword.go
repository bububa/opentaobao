package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryBidword 关键词报表查询
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
func TaobaoUniversalbpReportQueryBidword(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryBidwordAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryBidwordAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryBidwordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
