package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkcouponconvert 淘宝客-推广者-单品券高效转链
// taobao.tbk.coupon.convert
//
// 单品券高效转链API
func Taobaotbkcouponconvert(clt *core.SDKClient, req *tbk.TaobaotbkcouponconvertAPIRequest, session string) (*tbk.TaobaotbkcouponconvertAPIResponse, error) {
	var resp tbk.TaobaotbkcouponconvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
