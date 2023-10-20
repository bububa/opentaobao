package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallservicecentertpfundsrecoverquery 服务商资金权益逆向扣回的查询接口
// tmall.servicecenter.tp.funds.recover.query
//
// 服务商资金权益逆向扣回的查询接口
func Tmallservicecentertpfundsrecoverquery(clt *core.SDKClient, req *servicecenter.TmallservicecentertpfundsrecoverqueryAPIRequest, session string) (*servicecenter.TmallservicecentertpfundsrecoverqueryAPIResponse, error) {
	var resp servicecenter.TmallservicecentertpfundsrecoverqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
