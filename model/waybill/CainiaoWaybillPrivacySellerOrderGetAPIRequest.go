package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillPrivacySellerOrderGetAPIRequest
隐私面单商家订单查询 API请求
cainiao.waybill.privacy.seller.order.get

商家查询最近100天隐私面单记录 */
type CainiaoWaybillPrivacySellerOrderGetAPIRequest struct {
	model.Params
}

// New
