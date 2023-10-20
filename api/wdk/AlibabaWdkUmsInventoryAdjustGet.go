package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsinventoryadjustget 库调单-回流单
// alibaba.wdk.ums.inventory.adjust.get
//
// 库调单-回流单
func Alibabawdkumsinventoryadjustget(clt *core.SDKClient, req *wdk.AlibabawdkumsinventoryadjustgetAPIRequest, session string) (*wdk.AlibabawdkumsinventoryadjustgetAPIResponse, error) {
	var resp wdk.AlibabawdkumsinventoryadjustgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
