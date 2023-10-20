package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdate 新房楼栋修改e码
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
func AlibabaAlihouseNewhomeProjectBuildingEcodeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
