package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmerchantcategoryquery 三江erp对接类目查询接口
// alibaba.wdk.merchant.category.query
//
// 三江erp对接类目查询接口
func Alibabawdkmerchantcategoryquery(clt *core.SDKClient, req *wdk.AlibabawdkmerchantcategoryqueryAPIRequest, session string) (*wdk.AlibabawdkmerchantcategoryqueryAPIResponse, error) {
	var resp wdk.AlibabawdkmerchantcategoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
