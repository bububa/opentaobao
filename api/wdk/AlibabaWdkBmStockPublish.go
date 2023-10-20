package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkbmstockpublish 品牌营销涉及到的商品的库存同步接口
// alibaba.wdk.bm.stock.publish
//
// 用于操作sku的库存
func Alibabawdkbmstockpublish(clt *core.SDKClient, req *wdk.AlibabawdkbmstockpublishAPIRequest, session string) (*wdk.AlibabawdkbmstockpublishAPIResponse, error) {
	var resp wdk.AlibabawdkbmstockpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
