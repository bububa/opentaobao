package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitembrandquery 品牌信息查询
// alibaba.wdk.item.brand.query
//
// 品牌信息查询
func Alibabawdkitembrandquery(clt *core.SDKClient, req *wdkitem.AlibabawdkitembrandqueryAPIRequest, session string) (*wdkitem.AlibabawdkitembrandqueryAPIResponse, error) {
	var resp wdkitem.AlibabawdkitembrandqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
