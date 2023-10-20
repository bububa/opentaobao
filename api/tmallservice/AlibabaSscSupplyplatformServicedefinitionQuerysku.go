package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformservicedefinitionquerysku 服务sku查询
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
func Alibabasscsupplyplatformservicedefinitionquerysku(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformservicedefinitionqueryskuAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformservicedefinitionqueryskuAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformservicedefinitionqueryskuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
