package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceabilitydelete 删除服务能力
// alibaba.ssc.supplyplatform.serviceability.delete
//
// 删除服务能力
func Alibabasscsupplyplatformserviceabilitydelete(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceabilitydeleteAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceabilitydeleteAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceabilitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
