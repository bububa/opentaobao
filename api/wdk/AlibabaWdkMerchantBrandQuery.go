package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmerchantbrandquery 品牌查询接口
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
func Alibabawdkmerchantbrandquery(clt *core.SDKClient, req *wdk.AlibabawdkmerchantbrandqueryAPIRequest, session string) (*wdk.AlibabawdkmerchantbrandqueryAPIResponse, error) {
	var resp wdk.AlibabawdkmerchantbrandqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
