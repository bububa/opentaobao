package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalstoreauditquery ISV查询门店审核状态
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
func Alibabaalihealthdentalstoreauditquery(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalstoreauditqueryAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalstoreauditqueryAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalstoreauditqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
