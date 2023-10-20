package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingbuygiftactivityupdate 更新单品买赠活动
// alibaba.retail.marketing.buygift.activity.update
//
// 同城零售单品买赠活动更新
func Alibabaretailmarketingbuygiftactivityupdate(clt *core.SDKClient, req *wdk.AlibabaretailmarketingbuygiftactivityupdateAPIRequest, session string) (*wdk.AlibabaretailmarketingbuygiftactivityupdateAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingbuygiftactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
