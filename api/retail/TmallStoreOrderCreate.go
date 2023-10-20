package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// Tmallstoreordercreate 门店订单创建api
// tmall.store.order.create
//
// 门店订单创建api
func Tmallstoreordercreate(clt *core.SDKClient, req *retail.TmallstoreordercreateAPIRequest, session string) (*retail.TmallstoreordercreateAPIResponse, error) {
	var resp retail.TmallstoreordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
