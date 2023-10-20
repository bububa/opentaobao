package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmerchantskuquery 商家商品信息查询
// alibaba.wdk.item.merchantsku.query
//
// 商家商品信息查询
func Alibabawdkitemmerchantskuquery(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmerchantskuqueryAPIRequest, session string) (*wdkitem.AlibabawdkitemmerchantskuqueryAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmerchantskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
