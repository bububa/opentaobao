package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicesettlementfbbilldetailquery 服务商工单结算对账查询-流水查询
// tmall.service.settlement.fb.bill.detail.query
//
// 服务商工单结算对账查询-流水查询，用于查询服务工单费用流水，含服务费、退款、分成、提现等。
func Tmallservicesettlementfbbilldetailquery(clt *core.SDKClient, req *tmallservice.TmallservicesettlementfbbilldetailqueryAPIRequest, session string) (*tmallservice.TmallservicesettlementfbbilldetailqueryAPIResponse, error) {
	var resp tmallservice.TmallservicesettlementfbbilldetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
