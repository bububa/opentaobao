package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkitemstoreskuupdate 五道口商品中心门店商品修改
// alibaba.wdk.item.storesku.update
//
// 五道口商品中心门店商品修改
func Alibabawdkitemstoreskuupdate(clt *core.SDKClient, req *wdk.AlibabawdkitemstoreskuupdateAPIRequest, session string) (*wdk.AlibabawdkitemstoreskuupdateAPIResponse, error) {
	var resp wdk.AlibabawdkitemstoreskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
