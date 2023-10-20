package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabapricepromotionactivityquery 查询盒马帮档期活动详情
// alibaba.price.promotion.activity.query
//
// 查询盒马帮档期活动详情
func Alibabapricepromotionactivityquery(clt *core.SDKClient, req *wdk.AlibabapricepromotionactivityqueryAPIRequest, session string) (*wdk.AlibabapricepromotionactivityqueryAPIResponse, error) {
	var resp wdk.AlibabapricepromotionactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
