package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingbuygiftactivitysave 【同城零售】单品买赠活动保存
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
func Alibabaretailmarketingbuygiftactivitysave(clt *core.SDKClient, req *wdk.AlibabaretailmarketingbuygiftactivitysaveAPIRequest, session string) (*wdk.AlibabaretailmarketingbuygiftactivitysaveAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingbuygiftactivitysaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
