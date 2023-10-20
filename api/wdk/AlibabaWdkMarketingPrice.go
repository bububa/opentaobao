package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingprice 促销价签服务
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
func Alibabawdkmarketingprice(clt *core.SDKClient, req *wdk.AlibabawdkmarketingpriceAPIRequest, session string) (*wdk.AlibabawdkmarketingpriceAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
