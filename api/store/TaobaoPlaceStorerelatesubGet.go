package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestorerelatesubget 门店和子门店关系查找
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
func Taobaoplacestorerelatesubget(clt *core.SDKClient, req *store.TaobaoplacestorerelatesubgetAPIRequest, session string) (*store.TaobaoplacestorerelatesubgetAPIResponse, error) {
	var resp store.TaobaoplacestorerelatesubgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
