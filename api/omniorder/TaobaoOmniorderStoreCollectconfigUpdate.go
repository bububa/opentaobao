package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreCollectconfigUpdate 门店自提配置修改
// taobao.omniorder.store.collectconfig.update
//
// 修改门店自提配置内容
func TaobaoOmniorderStoreCollectconfigUpdate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreCollectconfigUpdateAPIRequest, resp *omniorder.TaobaoOmniorderStoreCollectconfigUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
