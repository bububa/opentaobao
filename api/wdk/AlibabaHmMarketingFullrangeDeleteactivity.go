package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeDeleteactivity 全场活动删除活动接口
// alibaba.hm.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
func AlibabaHmMarketingFullrangeDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeDeleteactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingFullrangeDeleteactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingFullrangeDeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
