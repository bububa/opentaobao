package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemdistributionbatchcancel 取消商品分销
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
func Alibabadchainaoxiangitemdistributionbatchcancel(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemdistributionbatchcancelAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemdistributionbatchcancelAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemdistributionbatchcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
