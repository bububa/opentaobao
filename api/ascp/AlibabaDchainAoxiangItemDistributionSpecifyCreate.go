package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemdistributionspecifycreate 指定分销商进行铺货(专享)
// alibaba.dchain.aoxiang.item.distribution.specify.create
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
func Alibabadchainaoxiangitemdistributionspecifycreate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemdistributionspecifycreateAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemdistributionspecifycreateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemdistributionspecifycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
