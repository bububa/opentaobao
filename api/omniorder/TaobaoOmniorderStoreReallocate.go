package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreReallocate rellocate
// taobao.omniorder.store.reallocate
//
// 门店发货提供改派接口
func TaobaoOmniorderStoreReallocate(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreReallocateAPIRequest, resp *omniorder.TaobaoOmniorderStoreReallocateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
