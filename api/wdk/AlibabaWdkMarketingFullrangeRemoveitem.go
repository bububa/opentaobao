package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingfullrangeremoveitem 全场活动删除购品
// alibaba.wdk.marketing.fullrange.removeitem
//
// 删除换购商品
func Alibabawdkmarketingfullrangeremoveitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingfullrangeremoveitemAPIRequest, session string) (*wdk.AlibabawdkmarketingfullrangeremoveitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingfullrangeremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
