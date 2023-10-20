package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmerchantskucreate 商家商品信息新建
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
func Alibabawdkitemmerchantskucreate(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmerchantskucreateAPIRequest, session string) (*wdkitem.AlibabawdkitemmerchantskucreateAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmerchantskucreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
