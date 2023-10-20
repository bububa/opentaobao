package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceworkerquitservicestore 工人退出网点
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
func Alibabasscsupplyplatformserviceworkerquitservicestore(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceworkerquitservicestoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
