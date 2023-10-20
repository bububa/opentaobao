package xhotelofficial

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// TaobaoXhotelOrderOfficialSettlePut 官网信用住结账接口
// taobao.xhotel.order.official.settle.put
//
// 用于酒店官网信用住商家结账调用
func TaobaoXhotelOrderOfficialSettlePut(clt *core.SDKClient, req *xhotelofficial.TaobaoXhotelOrderOfficialSettlePutAPIRequest, resp *xhotelofficial.TaobaoXhotelOrderOfficialSettlePutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
