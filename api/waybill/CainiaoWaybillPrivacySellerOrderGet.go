package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

/* CainiaoWaybillPrivacySellerOrderGet
隐私面单商家订单查询
cainiao.waybill.privacy.seller.order.get

商家查询最近100天隐私面单记录 */
func CainiaoWaybillPrivacySellerOrderGet(clt *core.SDKClient, req *waybill.CainiaoWaybillPrivacySellerOrderGetAPIRequest, session string) (*waybill.CainiaoWaybillPrivacySellerOrderGetAPIResponse, error) {
	var resp waybill.CainiaoWaybillPrivacySellerOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
