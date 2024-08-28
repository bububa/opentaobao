package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclUpdategrantroletouser 修改用户到角色关系
// alibaba.campus.acl.updategrantroletouser
//
// 修改用户到角色关系
func AlibabaCampusAclUpdategrantroletouser(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclUpdategrantroletouserAPIRequest, resp *campus.AlibabaCampusAclUpdategrantroletouserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
