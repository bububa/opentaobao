package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectBuilding 楼栋同步
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
func AlibabaAlihouseNewhomeProjectBuilding(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectBuildingAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectBuildingAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
