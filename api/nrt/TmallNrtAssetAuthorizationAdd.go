package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtAssetAuthorizationAdd 增加数据权限授权
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
func TmallNrtAssetAuthorizationAdd(clt *core.SDKClient, req *nrt.TmallNrtAssetAuthorizationAddAPIRequest, resp *nrt.TmallNrtAssetAuthorizationAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
