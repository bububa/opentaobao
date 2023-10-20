package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtstorerelationquery 喵零门店关系查询
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
func Tmallnrtstorerelationquery(clt *core.SDKClient, req *nrt.TmallnrtstorerelationqueryAPIRequest, session string) (*nrt.TmallnrtstorerelationqueryAPIResponse, error) {
	var resp nrt.TmallnrtstorerelationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
