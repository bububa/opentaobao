package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorpreturngoodsrefill 卖家回填物流信息
// taobao.rp.returngoods.refill
//
// 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
func Taobaorpreturngoodsrefill(clt *core.SDKClient, req *tbrefund.TaobaorpreturngoodsrefillAPIRequest, session string) (*tbrefund.TaobaorpreturngoodsrefillAPIResponse, error) {
	var resp tbrefund.TaobaorpreturngoodsrefillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
