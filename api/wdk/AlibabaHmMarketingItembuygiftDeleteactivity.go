package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitembuygiftdeleteactivity 删除买赠活动
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
func Alibabahmmarketingitembuygiftdeleteactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitembuygiftdeleteactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitembuygiftdeleteactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitembuygiftdeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
