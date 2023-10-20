package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorpreturngoodsrefuse 卖家拒绝退货
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
func Taobaorpreturngoodsrefuse(clt *core.SDKClient, req *tbrefund.TaobaorpreturngoodsrefuseAPIRequest, session string) (*tbrefund.TaobaorpreturngoodsrefuseAPIResponse, error) {
	var resp tbrefund.TaobaorpreturngoodsrefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
