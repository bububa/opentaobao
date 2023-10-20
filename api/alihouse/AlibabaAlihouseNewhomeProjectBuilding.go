package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectbuilding 楼栋同步
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
func Alibabaalihousenewhomeprojectbuilding(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectbuildingAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectbuildingAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectbuildingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
