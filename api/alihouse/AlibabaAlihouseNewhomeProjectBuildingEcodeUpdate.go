package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectbuildingecodeupdate 新房楼栋修改e码
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
func Alibabaalihousenewhomeprojectbuildingecodeupdate(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectbuildingecodeupdateAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
