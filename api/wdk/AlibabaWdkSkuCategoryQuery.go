package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucategoryquery 商家类目获取接口
// alibaba.wdk.sku.category.query
//
// 商家类目获取接口
func Alibabawdkskucategoryquery(clt *core.SDKClient, req *wdk.AlibabawdkskucategoryqueryAPIRequest, session string) (*wdk.AlibabawdkskucategoryqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskucategoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
