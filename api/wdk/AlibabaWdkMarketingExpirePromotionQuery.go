package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingexpirepromotionquery 短保优惠查询
// alibaba.wdk.marketing.expire.promotion.query
//
// 短保优惠查询
func Alibabawdkmarketingexpirepromotionquery(clt *core.SDKClient, req *wdk.AlibabawdkmarketingexpirepromotionqueryAPIRequest, session string) (*wdk.AlibabawdkmarketingexpirepromotionqueryAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingexpirepromotionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
