package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiQueryByTradecode 通过订单号查询电子面单通接口
// cainiao.waybill.ii.query.by.tradecode
//
// 通过订单号查看面单的信息
func CainiaoWaybillIiQueryByTradecode(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiQueryByTradecodeAPIRequest, resp *waybill.CainiaoWaybillIiQueryByTradecodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
