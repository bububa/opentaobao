package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitembuygiftcreateactivity 创建买赠活动
// alibaba.hm.marketing.itembuygift.createactivity
//
// 创建买赠活动
func Alibabahmmarketingitembuygiftcreateactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitembuygiftcreateactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitembuygiftcreateactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitembuygiftcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
