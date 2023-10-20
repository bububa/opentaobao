package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniitemskuget 获取全渠道门店商品sku
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
func Taobaoomniitemskuget(clt *core.SDKClient, req *omniorder.TaobaoomniitemskugetAPIRequest, session string) (*omniorder.TaobaoomniitemskugetAPIResponse, error) {
	var resp omniorder.TaobaoomniitemskugetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
