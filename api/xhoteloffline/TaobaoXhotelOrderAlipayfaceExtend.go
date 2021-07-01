package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

/* TaobaoXhotelOrderAlipayfaceExtend
信用住订单延住接口
taobao.xhotel.order.alipayface.extend

信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新 */
func TaobaoXhotelOrderAlipayfaceExtend(clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderAlipayfaceExtendAPIRequest, session string) (*xhoteloffline.TaobaoXhotelOrderAlipayfaceExtendAPIResponse, error) {
	var resp xhoteloffline.TaobaoXhotelOrderAlipayfaceExtendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
