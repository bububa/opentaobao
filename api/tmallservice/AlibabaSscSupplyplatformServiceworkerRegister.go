package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceworkerregister 服务商添加工人
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
func Alibabasscsupplyplatformserviceworkerregister(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceworkerregisterAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceworkerregisterAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceworkerregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
