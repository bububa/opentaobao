package rail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailIrCarrierGet 国际火车票铁路承运公司查询
// alitrip.rail.ir.carrier.get
//
// 国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
func AlitripRailIrCarrierGet(ctx context.Context, clt *core.SDKClient, req *rail.AlitripRailIrCarrierGetAPIRequest, resp *rail.AlitripRailIrCarrierGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
