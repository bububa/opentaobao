package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceAttrSetattr 新增业务属性实例接口
// alibaba.campus.space.attr.setattr
//
// 新增业务属性实例接口
func AlibabaCampusSpaceAttrSetattr(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceAttrSetattrAPIRequest, resp *campus.AlibabaCampusSpaceAttrSetattrAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
