package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingbuygiftactivitydelete 删除单品买赠活动
// alibaba.retail.marketing.buygift.activity.delete
//
// 同城零售单品特价活动删除
func Alibabaretailmarketingbuygiftactivitydelete(clt *core.SDKClient, req *wdk.AlibabaretailmarketingbuygiftactivitydeleteAPIRequest, session string) (*wdk.AlibabaretailmarketingbuygiftactivitydeleteAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingbuygiftactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
