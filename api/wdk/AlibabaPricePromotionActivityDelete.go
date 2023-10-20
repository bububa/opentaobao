package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabapricepromotionactivitydelete 删除档期活动
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
func Alibabapricepromotionactivitydelete(clt *core.SDKClient, req *wdk.AlibabapricepromotionactivitydeleteAPIRequest, session string) (*wdk.AlibabapricepromotionactivitydeleteAPIResponse, error) {
	var resp wdk.AlibabapricepromotionactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
