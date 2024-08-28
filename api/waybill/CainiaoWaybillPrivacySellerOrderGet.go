package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoWaybillPrivacySellerOrderGet 隐私面单商家订单查询
// cainiao.waybill.privacy.seller.order.get
//
// 商家查询最近100天隐私面单记录
func CainiaoWaybillPrivacySellerOrderGet(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoWaybillPrivacySellerOrderGetAPIRequest, resp *waybill.CainiaoWaybillPrivacySellerOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
