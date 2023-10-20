package xhotelofficial

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// Taobaoxhotelorderofficialcancel 官网信用住订单取消
// taobao.xhotel.order.official.cancel
//
// 官网信用住订单取消
func Taobaoxhotelorderofficialcancel(clt *core.SDKClient, req *xhotelofficial.TaobaoxhotelorderofficialcancelAPIRequest, session string) (*xhotelofficial.TaobaoxhotelorderofficialcancelAPIResponse, error) {
	var resp xhotelofficial.TaobaoxhotelorderofficialcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
