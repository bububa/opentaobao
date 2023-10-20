package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmerchantstoreskuupdate 门店商品信息修改
// alibaba.wdk.item.merchantstoresku.update
//
// 门店商品信息修改
func Alibabawdkitemmerchantstoreskuupdate(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmerchantstoreskuupdateAPIRequest, session string) (*wdkitem.AlibabawdkitemmerchantstoreskuupdateAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmerchantstoreskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
