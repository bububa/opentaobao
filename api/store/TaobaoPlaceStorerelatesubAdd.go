package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestorerelatesubadd 门店和子门店关系新增
// taobao.place.storerelatesub.add
//
// 门店和子门店关系新增
func Taobaoplacestorerelatesubadd(clt *core.SDKClient, req *store.TaobaoplacestorerelatesubaddAPIRequest, session string) (*store.TaobaoplacestorerelatesubaddAPIResponse, error) {
	var resp store.TaobaoplacestorerelatesubaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
