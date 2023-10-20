package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingexpirepromotionquery 短保优惠查询
// alibaba.hm.marketing.expire.promotion.query
//
// 短保优惠查询
func Alibabahmmarketingexpirepromotionquery(clt *core.SDKClient, req *wdk.AlibabahmmarketingexpirepromotionqueryAPIRequest, session string) (*wdk.AlibabahmmarketingexpirepromotionqueryAPIResponse, error) {
	var resp wdk.AlibabahmmarketingexpirepromotionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
