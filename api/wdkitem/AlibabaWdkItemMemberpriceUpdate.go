package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmemberpriceupdate 商品售价会员价修改
// alibaba.wdk.item.memberprice.update
//
// 商品售价会员价修改
func Alibabawdkitemmemberpriceupdate(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmemberpriceupdateAPIRequest, session string) (*wdkitem.AlibabawdkitemmemberpriceupdateAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmemberpriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
