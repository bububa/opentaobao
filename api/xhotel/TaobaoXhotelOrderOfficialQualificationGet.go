package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelOrderOfficialQualificationGet 官网信用住用户资质校验
// taobao.xhotel.order.official.qualification.get
//
// 官网信用住在下单前对用户进行资质校验，资质校验通过才能进行信用支付
func TaobaoXhotelOrderOfficialQualificationGet(clt *core.SDKClient, req *xhotel.TaobaoXhotelOrderOfficialQualificationGetAPIRequest, resp *xhotel.TaobaoXhotelOrderOfficialQualificationGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
