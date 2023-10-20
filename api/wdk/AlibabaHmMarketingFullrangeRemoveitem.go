package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingfullrangeremoveitem 全场活动删除购品
// alibaba.hm.marketing.fullrange.removeitem
//
// 删除换购商品
func Alibabahmmarketingfullrangeremoveitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingfullrangeremoveitemAPIRequest, session string) (*wdk.AlibabahmmarketingfullrangeremoveitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingfullrangeremoveitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
