package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtassetauthorizationdelete 移除资产数据权限授权关系
// tmall.nrt.asset.authorization.delete
//
// 移除资产数据权限授权关系
func Tmallnrtassetauthorizationdelete(clt *core.SDKClient, req *nrt.TmallnrtassetauthorizationdeleteAPIRequest, session string) (*nrt.TmallnrtassetauthorizationdeleteAPIResponse, error) {
	var resp nrt.TmallnrtassetauthorizationdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
