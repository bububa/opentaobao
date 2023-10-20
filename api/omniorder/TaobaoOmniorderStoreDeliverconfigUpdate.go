package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreDeliverconfigUpdate 修改门店发货配置内容
// taobao.omniorder.store.deliverconfig.update
//
// 修改门店发货配置内容
func TaobaoOmniorderStoreDeliverconfigUpdate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreDeliverconfigUpdateAPIRequest, resp *omniorder.TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
