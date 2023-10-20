package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucategoryadd 商家类目新增接口
// alibaba.wdk.sku.category.add
//
// 商家类目新增接口
func Alibabawdkskucategoryadd(clt *core.SDKClient, req *wdk.AlibabawdkskucategoryaddAPIRequest, session string) (*wdk.AlibabawdkskucategoryaddAPIResponse, error) {
	var resp wdk.AlibabawdkskucategoryaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
