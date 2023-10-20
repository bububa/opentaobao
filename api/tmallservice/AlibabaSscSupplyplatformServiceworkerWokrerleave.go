package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceworkerwokrerleave 工人请假
// alibaba.ssc.supplyplatform.serviceworker.wokrerleave
//
// 工人请假
func Alibabasscsupplyplatformserviceworkerwokrerleave(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceworkerwokrerleaveAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceworkerwokrerleaveAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceworkerwokrerleaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
