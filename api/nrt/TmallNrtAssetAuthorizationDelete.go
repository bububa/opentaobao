package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtAssetAuthorizationDelete 移除资产数据权限授权关系
// tmall.nrt.asset.authorization.delete
//
// 移除资产数据权限授权关系
func TmallNrtAssetAuthorizationDelete(clt *core.SDKClient, req *nrt.TmallNrtAssetAuthorizationDeleteAPIRequest, resp *nrt.TmallNrtAssetAuthorizationDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
