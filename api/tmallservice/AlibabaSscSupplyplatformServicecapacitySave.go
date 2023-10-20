package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformservicecapacitysave 保存服务容量
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
func Alibabasscsupplyplatformservicecapacitysave(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformservicecapacitysaveAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformservicecapacitysaveAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformservicecapacitysaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
