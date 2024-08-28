package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiServicestatusGet 供应商服务开通状态
// alibaba.happytrip.taxi.servicestatus.get
//
// 获取服务供应商在每个地区的服务开通状态、支持的车型等
func AlibabaHappytripTaxiServicestatusGet(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiServicestatusGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiServicestatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
