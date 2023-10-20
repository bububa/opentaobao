package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStoreExtendAdd 新增门店扩展属性
// taobao.place.store.extend.add
//
// 新增授权用户的门店扩展属性
func TaobaoPlaceStoreExtendAdd(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreExtendAddAPIRequest, resp *alsc.TaobaoPlaceStoreExtendAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
