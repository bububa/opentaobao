package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaSscPurchaseServicedefinitionParamQuery 查询采购服务定义参数信息
// alibaba.ssc.purchase.servicedefinition.param.query
//
// 查询采购服务定义参数信息
func AlibabaSscPurchaseServicedefinitionParamQuery(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest, resp *tmallsc.AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
