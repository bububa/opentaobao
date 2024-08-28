package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscServicecenterServicestoreQuery 根据天猫id查询门店信息
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
func AlibabaSscServicecenterServicestoreQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscServicecenterServicestoreQueryAPIRequest, resp *tmallservice.AlibabaSscServicecenterServicestoreQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
