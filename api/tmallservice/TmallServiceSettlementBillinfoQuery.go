package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettlementBillinfoQuery 查询工单结算信息
// tmall.service.settlement.billinfo.query
//
// 提供给服务商查询工单结算信息，包含结算的分成金额以及结算的收款明细，平台抽佣比例。用于服务商进行对账
func TmallServiceSettlementBillinfoQuery(clt *core.SDKClient, req *tmallservice.TmallServiceSettlementBillinfoQueryAPIRequest, session string) (*tmallservice.TmallServiceSettlementBillinfoQueryAPIResponse, error) {
	var resp tmallservice.TmallServiceSettlementBillinfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
