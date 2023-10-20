package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestorerelatesubdelete 门店和子门店关系删除
// taobao.place.storerelatesub.delete
//
// 门店和子门店关系删除
func Taobaoplacestorerelatesubdelete(clt *core.SDKClient, req *store.TaobaoplacestorerelatesubdeleteAPIRequest, session string) (*store.TaobaoplacestorerelatesubdeleteAPIResponse, error) {
	var resp store.TaobaoplacestorerelatesubdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
