package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformservicestoreoffline 网点下线
// alibaba.ssc.supplyplatform.servicestore.offline
//
// 网点下线功能
func Alibabasscsupplyplatformservicestoreoffline(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformservicestoreofflineAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformservicestoreofflineAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformservicestoreofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
