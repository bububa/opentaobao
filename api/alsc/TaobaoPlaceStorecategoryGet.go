package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStorecategoryGet 获取门店类目信息
// taobao.place.storecategory.get
//
// 获取门店类目信息
func TaobaoPlaceStorecategoryGet(clt *core.SDKClient, req *alsc.TaobaoPlaceStorecategoryGetAPIRequest, resp *alsc.TaobaoPlaceStorecategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
