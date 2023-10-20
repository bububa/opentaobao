package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtAssetAuthorizationAdd 增加数据权限授权
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
func TmallNrtAssetAuthorizationAdd(clt *core.SDKClient, req *nrt.TmallNrtAssetAuthorizationAddAPIRequest, session string) (*nrt.TmallNrtAssetAuthorizationAddAPIResponse, error) {
	var resp nrt.TmallNrtAssetAuthorizationAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
