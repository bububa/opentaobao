package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdate 新房楼栋修改e码
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
func AlibabaAlihouseNewhomeProjectBuildingEcodeUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
