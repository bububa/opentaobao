package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageAuthInfoGet 淘宝末端包裹信息获取
// taobao.top.package.auth.info.get
//
// 淘宝末端包裹信息获取
func TaobaoTopPackageAuthInfoGet(clt *core.SDKClient, req *topoaid.TaobaoTopPackageAuthInfoGetAPIRequest, resp *topoaid.TaobaoTopPackageAuthInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
