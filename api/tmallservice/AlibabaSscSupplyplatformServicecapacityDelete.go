package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformservicecapacitydelete 服务容量删除
// alibaba.ssc.supplyplatform.servicecapacity.delete
//
// 服务容量删除
func Alibabasscsupplyplatformservicecapacitydelete(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformservicecapacitydeleteAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformservicecapacitydeleteAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformservicecapacitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
