package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitembuygiftqueryactivity 查询买赠活动
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
func Alibabahmmarketingitembuygiftqueryactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitembuygiftqueryactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitembuygiftqueryactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitembuygiftqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
