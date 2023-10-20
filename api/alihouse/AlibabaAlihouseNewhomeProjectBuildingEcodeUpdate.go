package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdate 新房楼栋修改e码
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
func AlibabaAlihouseNewhomeProjectBuildingEcodeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
