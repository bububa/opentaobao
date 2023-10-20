package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Taobaoalitriptravelproductskuoverride （供销）产品级别日历价格库存修改，全量覆盖
// taobao.alitrip.travel.product.sku.override
//
// （供销）产品级别日历价格库存修改，全量覆盖
func Taobaoalitriptravelproductskuoverride(clt *core.SDKClient, req *travel.TaobaoalitriptravelproductskuoverrideAPIRequest, session string) (*travel.TaobaoalitriptravelproductskuoverrideAPIResponse, error) {
	var resp travel.TaobaoalitriptravelproductskuoverrideAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
