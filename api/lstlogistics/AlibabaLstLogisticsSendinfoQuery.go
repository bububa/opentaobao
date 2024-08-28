package lstlogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstLogisticsSendinfoQuery 供应商-异云-查询主订单包含的物流单
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
func AlibabaLstLogisticsSendinfoQuery(ctx context.Context, clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsSendinfoQueryAPIRequest, resp *lstlogistics.AlibabaLstLogisticsSendinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
