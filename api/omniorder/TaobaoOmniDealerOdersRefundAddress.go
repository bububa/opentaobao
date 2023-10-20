package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomnidealerodersrefundaddress 经销商查询逆向退货地址
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
func Taobaoomnidealerodersrefundaddress(clt *core.SDKClient, req *omniorder.TaobaoomnidealerodersrefundaddressAPIRequest, session string) (*omniorder.TaobaoomnidealerodersrefundaddressAPIResponse, error) {
	var resp omniorder.TaobaoomnidealerodersrefundaddressAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
