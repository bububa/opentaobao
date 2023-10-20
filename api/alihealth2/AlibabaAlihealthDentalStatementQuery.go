package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalstatementquery ISV查询对账单
// alibaba.alihealth.dental.statement.query
//
// ISV查询对账单
func Alibabaalihealthdentalstatementquery(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalstatementqueryAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalstatementqueryAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalstatementqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
