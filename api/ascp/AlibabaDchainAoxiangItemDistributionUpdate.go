package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemdistributionupdate 更新商品分销内容
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
func Alibabadchainaoxiangitemdistributionupdate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemdistributionupdateAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemdistributionupdateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemdistributionupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
