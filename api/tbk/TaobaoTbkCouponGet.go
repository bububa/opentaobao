package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkcouponget 淘宝客-公用-阿里妈妈推广券详情查询
// taobao.tbk.coupon.get
//
// 传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
func Taobaotbkcouponget(clt *core.SDKClient, req *tbk.TaobaotbkcoupongetAPIRequest, session string) (*tbk.TaobaotbkcoupongetAPIResponse, error) {
	var resp tbk.TaobaotbkcoupongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
