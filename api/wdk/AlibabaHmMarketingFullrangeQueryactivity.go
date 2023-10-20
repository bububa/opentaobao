package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingfullrangequeryactivity 全场活动查询活动
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
func Alibabahmmarketingfullrangequeryactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingfullrangequeryactivityAPIRequest, session string) (*wdk.AlibabahmmarketingfullrangequeryactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingfullrangequeryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
