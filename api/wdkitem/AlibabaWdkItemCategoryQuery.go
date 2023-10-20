package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemcategoryquery 类目查询接口
// alibaba.wdk.item.category.query
//
// 类目查询接口
func Alibabawdkitemcategoryquery(clt *core.SDKClient, req *wdkitem.AlibabawdkitemcategoryqueryAPIRequest, session string) (*wdkitem.AlibabawdkitemcategoryqueryAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemcategoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
