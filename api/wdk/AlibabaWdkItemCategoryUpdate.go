package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkitemcategoryupdate 修改类目
// alibaba.wdk.item.category.update
//
// 修改类目
func Alibabawdkitemcategoryupdate(clt *core.SDKClient, req *wdk.AlibabawdkitemcategoryupdateAPIRequest, session string) (*wdk.AlibabawdkitemcategoryupdateAPIResponse, error) {
	var resp wdk.AlibabawdkitemcategoryupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
