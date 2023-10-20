package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// Taobaoxhotelorderalipayfacecheck 线下信用住买家资格校验接口
// taobao.xhotel.order.alipayface.check
//
// 接口用于校验买家是否具有使用酒店信用住的资格
func Taobaoxhotelorderalipayfacecheck(clt *core.SDKClient, req *xhoteloffline.TaobaoxhotelorderalipayfacecheckAPIRequest, session string) (*xhoteloffline.TaobaoxhotelorderalipayfacecheckAPIResponse, error) {
	var resp xhoteloffline.TaobaoxhotelorderalipayfacecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
