package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurcmallgoodsstatussync 第三方商城接入采购商城-商品状态同步
// alibaba.pur.cmall.goods.status.sync
//
// 第三方商城接入采购商城-商品状态同步
func Alibabapurcmallgoodsstatussync(clt *core.SDKClient, req *pur.AlibabapurcmallgoodsstatussyncAPIRequest, session string) (*pur.AlibabapurcmallgoodsstatussyncAPIResponse, error) {
	var resp pur.AlibabapurcmallgoodsstatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
