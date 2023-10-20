package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsinventorycheckget 盘点结果单-回流单
// alibaba.wdk.ums.inventory.check.get
//
// 盘点结果单-回流单
func Alibabawdkumsinventorycheckget(clt *core.SDKClient, req *wdk.AlibabawdkumsinventorycheckgetAPIRequest, session string) (*wdk.AlibabawdkumsinventorycheckgetAPIResponse, error) {
	var resp wdk.AlibabawdkumsinventorycheckgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
