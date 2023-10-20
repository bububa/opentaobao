package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangitemdistributiondescribe 分销商品文描
// alibaba.dchain.aoxiang.item.distribution.describe
//
// 分销商品文描
func Alibabadchainaoxiangitemdistributiondescribe(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangitemdistributiondescribeAPIRequest, session string) (*ascp.AlibabadchainaoxiangitemdistributiondescribeAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangitemdistributiondescribeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
