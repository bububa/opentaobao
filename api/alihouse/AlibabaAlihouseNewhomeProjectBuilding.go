package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectBuilding 楼栋同步
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
func AlibabaAlihouseNewhomeProjectBuilding(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectBuildingAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectBuildingAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectBuildingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
