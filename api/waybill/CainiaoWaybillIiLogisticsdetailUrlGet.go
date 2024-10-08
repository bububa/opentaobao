package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillIiLogisticsdetailUrlGet 电子面单物流详情授权url获取
// cainiao.waybill.ii.logisticsdetail.url.get
//
// 获取电子面单物流详情授权访问的H5 url
func CainiaoWaybillIiLogisticsdetailUrlGet(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest, resp *waybill.CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
