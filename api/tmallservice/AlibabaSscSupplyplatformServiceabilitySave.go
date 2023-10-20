package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceabilitysave 保存服务能力
// alibaba.ssc.supplyplatform.serviceability.save
//
// 保存服务能力
func Alibabasscsupplyplatformserviceabilitysave(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceabilitysaveAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceabilitysaveAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceabilitysaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
