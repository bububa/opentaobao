package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestorerelationquery 门店关系查询
// taobao.place.store.relation.query
//
// 查询门店关系
func Taobaoplacestorerelationquery(clt *core.SDKClient, req *store.TaobaoplacestorerelationqueryAPIRequest, session string) (*store.TaobaoplacestorerelationqueryAPIResponse, error) {
	var resp store.TaobaoplacestorerelationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
