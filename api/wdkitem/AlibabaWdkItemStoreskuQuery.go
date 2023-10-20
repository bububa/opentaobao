package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemstoreskuquery 门店商品信息查询
// alibaba.wdk.item.storesku.query
//
// 门店商品信息查询
func Alibabawdkitemstoreskuquery(clt *core.SDKClient, req *wdkitem.AlibabawdkitemstoreskuqueryAPIRequest, session string) (*wdkitem.AlibabawdkitemstoreskuqueryAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemstoreskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
