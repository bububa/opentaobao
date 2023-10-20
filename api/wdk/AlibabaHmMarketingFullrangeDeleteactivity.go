package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingfullrangedeleteactivity 全场活动删除活动接口
// alibaba.hm.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
func Alibabahmmarketingfullrangedeleteactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingfullrangedeleteactivityAPIRequest, session string) (*wdk.AlibabahmmarketingfullrangedeleteactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingfullrangedeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
