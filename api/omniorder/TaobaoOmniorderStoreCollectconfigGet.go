package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreCollectconfigGet 查询门店自提配置内容
// taobao.omniorder.store.collectconfig.get
//
// 查询门店自提配置内容
func TaobaoOmniorderStoreCollectconfigGet(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreCollectconfigGetAPIRequest, resp *omniorder.TaobaoOmniorderStoreCollectconfigGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
