package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanpayresultquery 百川支付完成回调
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
func Taobaobaichuanpayresultquery(clt *core.SDKClient, req *baichuan.TaobaobaichuanpayresultqueryAPIRequest, session string) (*baichuan.TaobaobaichuanpayresultqueryAPIResponse, error) {
	var resp baichuan.TaobaobaichuanpayresultqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
