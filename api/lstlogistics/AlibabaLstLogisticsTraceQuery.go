package lstlogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstLogisticsTraceQuery 供应商-异云-查询运单物流追踪信息
// alibaba.lst.logistics.trace.query
//
// 查询LP单物流追踪信息
func AlibabaLstLogisticsTraceQuery(ctx context.Context, clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsTraceQueryAPIRequest, resp *lstlogistics.AlibabaLstLogisticsTraceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
