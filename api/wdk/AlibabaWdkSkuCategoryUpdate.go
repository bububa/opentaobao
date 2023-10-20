package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucategoryupdate 商家类目修改接口
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
func Alibabawdkskucategoryupdate(clt *core.SDKClient, req *wdk.AlibabawdkskucategoryupdateAPIRequest, session string) (*wdk.AlibabawdkskucategoryupdateAPIResponse, error) {
	var resp wdk.AlibabawdkskucategoryupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
