package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabaservicesettlementquery 服务平台结算单明细查询服务
// alibaba.service.settlement.query
//
// 给服务商提供结算单明细查询功能
func Alibabaservicesettlementquery(clt *core.SDKClient, req *tmallsc.AlibabaservicesettlementqueryAPIRequest, session string) (*tmallsc.AlibabaservicesettlementqueryAPIResponse, error) {
	var resp tmallsc.AlibabaservicesettlementqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
