package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetspaceunitlistwithattr 空间单元列表带业务属性实例
// alibaba.campus.space.unit.getspaceunitlistwithattr
//
// 空间单元列表带业务属性实例
func AlibabaCampusSpaceUnitGetspaceunitlistwithattr(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
