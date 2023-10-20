package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingfullrangecreateactivity 创建全场活动
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
func Alibabahmmarketingfullrangecreateactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingfullrangecreateactivityAPIRequest, session string) (*wdk.AlibabahmmarketingfullrangecreateactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingfullrangecreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
