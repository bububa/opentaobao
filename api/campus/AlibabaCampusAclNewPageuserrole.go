package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewPageuserrole 分页查询管理员
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
func AlibabaCampusAclNewPageuserrole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewPageuserroleAPIRequest, resp *campus.AlibabaCampusAclNewPageuserroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
