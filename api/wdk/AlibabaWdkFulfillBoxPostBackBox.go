package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillboxpostbackbox RT收箱回传
// alibaba.wdk.fulfill.box.post.back.box
//
// RT收箱后，信息同步履约，履约同通知UMS 容器管理
func Alibabawdkfulfillboxpostbackbox(clt *core.SDKClient, req *wdk.AlibabawdkfulfillboxpostbackboxAPIRequest, session string) (*wdk.AlibabawdkfulfillboxpostbackboxAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillboxpostbackboxAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
