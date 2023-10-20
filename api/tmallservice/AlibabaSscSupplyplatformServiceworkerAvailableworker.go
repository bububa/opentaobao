package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceworkeravailableworker 查询可用工人
// alibaba.ssc.supplyplatform.serviceworker.availableworker
//
// 可用工人查询
func Alibabasscsupplyplatformserviceworkeravailableworker(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceworkeravailableworkerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
