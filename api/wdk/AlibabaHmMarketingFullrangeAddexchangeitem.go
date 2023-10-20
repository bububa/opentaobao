package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeAddexchangeitem 全场增加换购品
// alibaba.hm.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
func AlibabaHmMarketingFullrangeAddexchangeitem(clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest, session string) (*wdk.AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
