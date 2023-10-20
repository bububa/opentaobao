package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscsupplyplatformserviceinventoryquery 服务库存查询
// alibaba.ssc.supplyplatform.service.inventory.query
//
// 查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。
func Alibabasscsupplyplatformserviceinventoryquery(clt *core.SDKClient, req *tmallservice.AlibabasscsupplyplatformserviceinventoryqueryAPIRequest, session string) (*tmallservice.AlibabasscsupplyplatformserviceinventoryqueryAPIResponse, error) {
	var resp tmallservice.AlibabasscsupplyplatformserviceinventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
