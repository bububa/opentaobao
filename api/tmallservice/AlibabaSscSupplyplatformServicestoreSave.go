package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformservicestoresave 保存网点
// alibaba.ssc.supplyplatform.servicestore.save
//
// 网点创建、修改
func Alibabasscsupplyplatformservicestoresave(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformservicestoresaveAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformservicestoresaveAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformservicestoresaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
