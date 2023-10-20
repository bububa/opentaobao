package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtstorecontractquery 摊位合同查询接口
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
func Tmallnrtstorecontractquery(clt *core.SDKClient, req *nrt.TmallnrtstorecontractqueryAPIRequest, session string) (*nrt.TmallnrtstorecontractqueryAPIResponse, error) {
	var resp nrt.TmallnrtstorecontractqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
