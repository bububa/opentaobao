package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybillprivacysellerorderget 隐私面单商家订单查询
// cainiao.waybill.privacy.seller.order.get
//
// 商家查询最近100天隐私面单记录
func Cainiaowaybillprivacysellerorderget(clt *core.SDKClient, req *waybill.CainiaowaybillprivacysellerordergetAPIRequest, session string) (*waybill.CainiaowaybillprivacysellerordergetAPIResponse, error) {
	var resp waybill.CainiaowaybillprivacysellerordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
