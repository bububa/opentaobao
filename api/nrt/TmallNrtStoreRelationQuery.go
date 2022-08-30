package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreRelationQuery 喵零门店关系查询
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
func TmallNrtStoreRelationQuery(clt *core.SDKClient, req *nrt.TmallNrtStoreRelationQueryAPIRequest, session string) (*nrt.TmallNrtStoreRelationQueryAPIResponse, error) {
	var resp nrt.TmallNrtStoreRelationQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
