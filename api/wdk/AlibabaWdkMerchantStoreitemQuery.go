package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmerchantstoreitemquery 门店商品信心查询
// alibaba.wdk.merchant.storeitem.query
//
// 门店商品信心查询
func Alibabawdkmerchantstoreitemquery(clt *core.SDKClient, req *wdk.AlibabawdkmerchantstoreitemqueryAPIRequest, session string) (*wdk.AlibabawdkmerchantstoreitemqueryAPIResponse, error) {
	var resp wdk.AlibabawdkmerchantstoreitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
