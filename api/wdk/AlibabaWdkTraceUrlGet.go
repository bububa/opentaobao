package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdktraceurlget 溯源url透出
// alibaba.wdk.trace.url.get
//
// 根据shopId和skuCode返回商品溯源url
func Alibabawdktraceurlget(clt *core.SDKClient, req *wdk.AlibabawdktraceurlgetAPIRequest, session string) (*wdk.AlibabawdktraceurlgetAPIResponse, error) {
	var resp wdk.AlibabawdktraceurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
