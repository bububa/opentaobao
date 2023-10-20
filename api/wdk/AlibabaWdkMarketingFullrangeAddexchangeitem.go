package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingfullrangeaddexchangeitem 全场增加换购品
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
func Alibabawdkmarketingfullrangeaddexchangeitem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest, session string) (*wdk.AlibabawdkmarketingfullrangeaddexchangeitemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingfullrangeaddexchangeitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
