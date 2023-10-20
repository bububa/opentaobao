package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkitemtraceurlget 根据shopId和skuCode返回商品静态溯源url
// alibaba.wdk.item.trace.url.get
//
// 根据shopId和skuCode返回商品静态溯源url
func Alibabawdkitemtraceurlget(clt *core.SDKClient, req *wdk.AlibabawdkitemtraceurlgetAPIRequest, session string) (*wdk.AlibabawdkitemtraceurlgetAPIResponse, error) {
	var resp wdk.AlibabawdkitemtraceurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
