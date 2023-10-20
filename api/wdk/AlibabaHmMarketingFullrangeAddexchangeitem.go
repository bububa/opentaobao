package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingfullrangeaddexchangeitem 全场增加换购品
// alibaba.hm.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
func Alibabahmmarketingfullrangeaddexchangeitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingfullrangeaddexchangeitemAPIRequest, session string) (*wdk.AlibabahmmarketingfullrangeaddexchangeitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingfullrangeaddexchangeitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
