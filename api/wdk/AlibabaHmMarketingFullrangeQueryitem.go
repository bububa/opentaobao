package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingfullrangequeryitem 全场活动查询换购品
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
func Alibabahmmarketingfullrangequeryitem(clt *core.SDKClient, req *wdk.AlibabahmmarketingfullrangequeryitemAPIRequest, session string) (*wdk.AlibabahmmarketingfullrangequeryitemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingfullrangequeryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
