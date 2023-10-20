package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemdistributionspecifyupdate 指定分销商进行铺货(专享) - 修改
// alibaba.dchain.aoxiang.item.distribution.specify.update
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
func Alibabadchainaoxiangitemdistributionspecifyupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemdistributionspecifyupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemdistributionspecifyupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
