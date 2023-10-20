package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmerchantstoreskucreate 门店商品信息新建
// alibaba.wdk.item.merchantstoresku.create
//
// 门店商品信息新建
func Alibabawdkitemmerchantstoreskucreate(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmerchantstoreskucreateAPIRequest, session string) (*wdkitem.AlibabawdkitemmerchantstoreskucreateAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmerchantstoreskucreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
