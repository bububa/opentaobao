package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemstoreskustatusupdate 修改门店商品状态
// alibaba.wdk.item.storeskustatus.update
//
// 五道口商品 修改门店商品状态
func Alibabawdkitemstoreskustatusupdate(clt *core.SDKClient, req *wdkitem.AlibabawdkitemstoreskustatusupdateAPIRequest, session string) (*wdkitem.AlibabawdkitemstoreskustatusupdateAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemstoreskustatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
