package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtAssetAuthorizationDelete 移除资产数据权限授权关系
// tmall.nrt.asset.authorization.delete
//
// 移除资产数据权限授权关系
func TmallNrtAssetAuthorizationDelete(clt *core.SDKClient, req *nrt.TmallNrtAssetAuthorizationDeleteAPIRequest, session string) (*nrt.TmallNrtAssetAuthorizationDeleteAPIResponse, error) {
	var resp nrt.TmallNrtAssetAuthorizationDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
