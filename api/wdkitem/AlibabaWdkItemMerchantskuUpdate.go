package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmerchantskuupdate 商家商品修改
// alibaba.wdk.item.merchantsku.update
//
// 商家商品修改
func Alibabawdkitemmerchantskuupdate(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmerchantskuupdateAPIRequest, session string) (*wdkitem.AlibabawdkitemmerchantskuupdateAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmerchantskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
