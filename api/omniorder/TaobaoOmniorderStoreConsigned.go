package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreConsigned Pos端门店发货
// taobao.omniorder.store.consigned
//
// ISV Pos端门店发货，通知星盘
func TaobaoOmniorderStoreConsigned(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreConsignedAPIRequest, resp *omniorder.TaobaoOmniorderStoreConsignedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
