package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingbuygiftactivitycreate 创建买赠活动
// alibaba.retail.marketing.buygift.activity.create
//
// 同城供给买赠活动创建
func Alibabaretailmarketingbuygiftactivitycreate(clt *core.SDKClient, req *wdk.AlibabaretailmarketingbuygiftactivitycreateAPIRequest, session string) (*wdk.AlibabaretailmarketingbuygiftactivitycreateAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingbuygiftactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
