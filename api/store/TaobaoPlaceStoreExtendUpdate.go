package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoreExtendUpdate 商户门店拓展信息更新接口
// taobao.place.store.extend.update
//
// 更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
func TaobaoPlaceStoreExtendUpdate(clt *core.SDKClient, req *store.TaobaoPlaceStoreExtendUpdateAPIRequest, resp *store.TaobaoPlaceStoreExtendUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
