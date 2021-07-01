package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniorderStoreCollectconfigUpdate
门店自提配置修改
taobao.omniorder.store.collectconfig.update

修改门店自提配置内容 */
func TaobaoOmniorderStoreCollectconfigUpdate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreCollectconfigUpdateAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreCollectconfigUpdateAPIResponse, error) {
	var resp omniorder.TaobaoOmniorderStoreCollectconfigUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
