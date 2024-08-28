package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGrantpermiitemstouser 给人直接授权
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
func AlibabaCampusAclGrantpermiitemstouser(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclGrantpermiitemstouserAPIRequest, resp *campus.AlibabaCampusAclGrantpermiitemstouserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
