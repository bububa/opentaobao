package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsConsignModify 修改物流公司和运单号
// alibaba.ascp.logistics.consign.modify
//
// 修改物流公司和运单号
func AlibabaAscpLogisticsConsignModify(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsConsignModifyAPIRequest, resp *tblogistics.AlibabaAscpLogisticsConsignModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
