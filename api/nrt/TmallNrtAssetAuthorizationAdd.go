package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtassetauthorizationadd 增加数据权限授权
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
func Tmallnrtassetauthorizationadd(clt *core.SDKClient, req *nrt.TmallnrtassetauthorizationaddAPIRequest, session string) (*nrt.TmallnrtassetauthorizationaddAPIResponse, error) {
	var resp nrt.TmallnrtassetauthorizationaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
