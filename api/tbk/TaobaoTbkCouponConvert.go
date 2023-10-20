package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkCouponConvert 淘宝客-推广者-单品券高效转链
// taobao.tbk.coupon.convert
//
// 单品券高效转链API
func TaobaoTbkCouponConvert(clt *core.SDKClient, req *tbk.TaobaoTbkCouponConvertAPIRequest, session string) (*tbk.TaobaoTbkCouponConvertAPIResponse, error) {
	var resp tbk.TaobaoTbkCouponConvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
