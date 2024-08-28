package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetspacegroupwithattr 空间分组id查业务属性实例
// alibaba.campus.space.group.getspacegroupwithattr
//
// 空间分组id查业务属性实例
func AlibabaCampusSpaceGroupGetspacegroupwithattr(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest, resp *campus.AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
